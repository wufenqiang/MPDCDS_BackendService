package repo

import (
	"MPDCDS_BackendService/src/datasource"
	"MPDCDS_BackendService/src/logger"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"
	"go.uber.org/zap"
)

var ApiFileRepo = new(apiFileRepo)

type apiFileRepo struct{}

//根据索引名和标识查询文件信息
func (a *apiFileRepo) GetFileByIndexNameAndDirId(indexName, dirId string) (r []models.ApiFile) {
	if indexName == "" || dirId == "" {
		return
	}

	dataClient := datasource.ES.GetDB()

	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewTermQuery("dir_id.keyword", dirId))

	res, err := dataClient.Search(indexName).
		Size(10000).
		From(0).
		Sort("create_time", true).
		Query(boolQ).Do(context.Background())

	if err != nil {
		logger.GetLogger().Error("GetDirByPath", zap.Error(err))
		return
	}

	if res == nil {
		return
	}

	for _, item := range res.Hits.Hits {
		var apiFile models.ApiFile
		data, _ := item.Source.MarshalJSON()
		json.Unmarshal(data, &apiFile)
		apiFile.Id = item.Id
		r = append(r, apiFile)
	}
	return
}

//根据索引名和标识、文件名称查询文件真实地址
func (a *apiFileRepo) GetFileByIndexNameAndDirIdAndFileName(indexName, dirId, fileName string) (r models.ApiFile) {
	if indexName == "" || dirId == "" {
		return
	}
	esClient := datasource.ES.GetDB()

	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewTermQuery("dir_id.keyword", dirId))
	boolQ.Must(elastic.NewTermQuery("file_name.keyword", fileName))
	res, err := esClient.Search(indexName).Query(boolQ).Do(context.Background())
	if err != nil {
		logger.GetLogger().Error("GetFileByIndexNameAndDirIdAndFileName", zap.Error(err))
		return
	}

	if res == nil {
		return
	}
	var apiFile models.ApiFile
	for _, item := range res.Hits.Hits {
		data, _ := item.Source.MarshalJSON()
		json.Unmarshal(data, &apiFile)
		apiFile.Id = item.Id
		break
	}
	return apiFile
}
