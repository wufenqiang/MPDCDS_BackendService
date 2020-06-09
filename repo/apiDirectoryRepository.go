package repo

import (
	esdatasource "MPDCDS_BackendService/datasource/elasticsearch"
	"MPDCDS_BackendService/logger"
	"MPDCDS_BackendService/models"
	"MPDCDS_BackendService/utils"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type ApiDirectoryRepository interface {
	GetDirByParentPath(path string) []models.ApiDirectory
	GetDirByCurrentPath(currentPath string) models.ApiDirectory
}

func NewApiDirectoryRepository() ApiDirectoryRepository {
	return &apiDirectoryRepository{}
}

type apiDirectoryRepository struct{}

func (a apiDirectoryRepository) GetDirByParentPath(path string) (r []models.ApiDirectory) {
	if path == "" {
		return
	}
	esClient := esdatasource.GetESClient()

	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewTermQuery("parent_dir.keyword", path))
	res, err := esClient.Search(utils.UnMarshal(models.ApiDirectory{})).
		Size(10000).
		From(0).
		Query(boolQ).Do(context.Background())

	if err != nil {
		logger.GetLogger().Error("GetDirByPath", zap.Error(err))
		return
	}

	if res == nil {
		return
	}

	for _, item := range res.Hits.Hits {
		var apiDirectory models.ApiDirectory
		data, _ := item.Source.MarshalJSON()
		json.Unmarshal(data, &apiDirectory)
		apiDirectory.Id = item.Id
		r = append(r, apiDirectory)
	}

	return
}

func (a apiDirectoryRepository) GetDirByCurrentPath(currentPath string) (r models.ApiDirectory) {
	if currentPath == "" {
		return
	}
	esClient := esdatasource.GetESClient()

	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewTermQuery("current_dir.keyword", currentPath))

	res, err := esClient.Search(utils.UnMarshal(models.ApiDirectory{})).
		Size(1).
		From(0).
		Query(boolQ).Do(context.Background())

	if err != nil {
		logger.GetLogger().Error("GetDirByPath", zap.Error(err))
		return
	}

	if res == nil {
		return
	}

	for _, item := range res.Hits.Hits {
		data, _ := item.Source.MarshalJSON()
		json.Unmarshal(data, &r)
		r.Id = item.Id
	}
	return
}
