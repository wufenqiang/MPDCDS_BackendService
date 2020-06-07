package repo

import (
	esdatasource "MPDCDS_BackendService/datasource/elasticsearch"
	"MPDCDS_BackendService/logger"
	"MPDCDS_BackendService/models"
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"reflect"
)

type ApiFileRepository interface {
	//根据目录ID或者接入数据信息表ID获取文件信息
	GetFileByDirId(apiDirId, apiDataInfoId string) []models.ApiFile
}

func NewApiFileRepository() ApiFileRepository {
	return &apiFileRepository{}
}

type apiFileRepository struct{}

func (a apiFileRepository) GetFileByDirId(apiDirId, apiDataInfoId string) []models.ApiFile {
	esClient := esdatasource.GetESClient()

	boolQ := elastic.NewBoolQuery()
	if apiDirId != "" {
		boolQ.Must(elastic.NewQueryStringQuery("dir_id:" + apiDirId))
	}
	if apiDataInfoId != "" {
		boolQ.Must(elastic.NewQueryStringQuery("access_id:" + apiDataInfoId))
	}

	res, err := esClient.Search("api_file").
		Size(10000).
		From(0).
		Query(boolQ).Do(context.Background())

	if err != nil {
		logger.GetLogger().Error("GetDirByPath", zap.Error(err))
	}

	if res == nil {
		return []models.ApiFile{}
	}

	var apiFiles []models.ApiFile
	for _, item := range res.Each(reflect.TypeOf(apiFiles)) { //从搜索结果中取数据的方法
		apiFiles = append(apiFiles, item.(models.ApiFile))
	}
	return apiFiles
}
