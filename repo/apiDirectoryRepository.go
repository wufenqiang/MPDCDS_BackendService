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

type ApiDirectoryRepository interface {
	GetDirByPath(path string) models.ApiDirectory
}

func NewApiDirectoryRepository() ApiDirectoryRepository {
	return &apiDirectoryRepository{}
}

type apiDirectoryRepository struct{}

func (a apiDirectoryRepository) GetDirByPath(path string) models.ApiDirectory {
	esClient := esdatasource.GetESClient()

	q := elastic.NewQueryStringQuery("path:" + path)
	res, err := esClient.Search("api_directory").
		Size(1).
		From(0).
		Query(q).Do(context.Background())

	if err != nil {
		logger.GetLogger().Error("GetDirByPath", zap.Error(err))
	}

	if res == nil {
		return models.ApiDirectory{}
	}

	var apiDirectory models.ApiDirectory
	for _, item := range res.Each(reflect.TypeOf(apiDirectory)) { //从搜索结果中取数据的方法
		apiDirectory = item.(models.ApiDirectory)
	}
	return apiDirectory
}
