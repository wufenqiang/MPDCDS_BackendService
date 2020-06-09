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

type ApiDataInfoRepository interface {
	//根据用户ID获取订单
	GetApiDataInfoById(ids []string) []models.ApiDataInfo
}

func NewApiDataInfoRepository() ApiDataInfoRepository {
	return &apiDataInfoRepository{}
}

type apiDataInfoRepository struct{}

func (a apiDataInfoRepository) GetApiDataInfoById(ids []string) (r []models.ApiDataInfo) {
	esClient := esdatasource.GetESClient()

	if ids == nil || len(ids) < 1 {
		return
	}

	boolQ := elastic.NewBoolQuery()
	boolQ.Filter(elastic.NewIdsQuery().Ids(ids...))

	res, err := esClient.Search(utils.UnMarshal(models.ApiDataInfo{})).
		Size(10000).
		From(0).
		Query(boolQ).Do(context.Background())

	if err != nil {
		logger.GetLogger().Error("GetApiDataInfoById", zap.Error(err))
		return
	}

	if res == nil {
		return
	}

	for _, item := range res.Hits.Hits {
		var apiApiDataInfo models.ApiDataInfo
		data, _ := item.Source.MarshalJSON()
		json.Unmarshal(data, &apiApiDataInfo)
		apiApiDataInfo.Id = item.Id
		r = append(r, apiApiDataInfo)
	}
	return
}
