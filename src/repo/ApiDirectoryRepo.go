package repo

import (
	"MPDCDS_BackendService/src/datasource"
	"MPDCDS_BackendService/src/logger"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"

	proutils "gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/utils"

	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

var ApiDirectoryRepo = new(apiDirectoryRepo)

type apiDirectoryRepo struct{}

func (a *apiDirectoryRepo) GetDirByParentPath(path string) (r []models.ApiDirectory) {
	if path == "" {
		return
	}

	dataClient := datasource.ES.GetDB()

	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewTermQuery("parent_dir.keyword", path))
	res, err := dataClient.Search(proutils.UnMarshal(models.ApiDirectory{})).
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

func (a *apiDirectoryRepo) GetDirByCurrentPath(currentPath string) (r models.ApiDirectory) {
	if currentPath == "" {
		return
	}
	dataClient := datasource.ES.GetDB()

	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewTermQuery("current_dir.keyword", currentPath))

	res, err := dataClient.Search(proutils.UnMarshal(models.ApiDirectory{})).
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
