package repo

import (
	esdatasource "MPDCDS_BackendService/datasource/elasticsearch"
	"MPDCDS_BackendService/logger"
	"MPDCDS_BackendService/models"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type ApiFileRepository interface {
	//根据索引名和标识查询文件信息
	GetFileByIndexNameAndDirId(indexName, dirId string) []models.ApiFile
	//根据索引名和标识、文件名称查询文件真实地址
	GetFileByIndexNameAndDirIdAndFileName(indexName, dirId, fileName string) (r models.ApiFile)
}

func NewApiFileRepository() ApiFileRepository {
	return &apiFileRepository{}
}

type apiFileRepository struct{}

func (a apiFileRepository) GetFileByIndexNameAndDirId(indexName, dirId string) (r []models.ApiFile) {
	if indexName == "" || dirId == "" {
		return
	}
	esClient := esdatasource.GetESClient()

	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewTermQuery("dir_id.keyword", dirId))

	res, err := esClient.Search(indexName).
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

func (a apiFileRepository) GetFileByIndexNameAndDirIdAndFileName(indexName, dirId, fileName string) (r models.ApiFile) {
	if indexName == "" || dirId == "" {
		return
	}
	esClient := esdatasource.GetESClient()

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
