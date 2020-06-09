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
	"time"
)

type ApiOrderRepository interface {
	//根据用户ID获取订单
	GetOrderByUserId(userId string) []models.ApiOrder
}

func NewApiOrderRepository() ApiOrderRepository {
	return &apiOrderRepository{}
}

type apiOrderRepository struct{}

func (a apiOrderRepository) GetOrderByUserId(userId string) (r []models.ApiOrder) {
	if userId == "" {
		return
	}

	esClient := esdatasource.GetESClient()

	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewQueryStringQuery("user_id:" + userId))
	boolQ.Must(elastic.NewQueryStringQuery("status:3"))
	boolQ.Must(elastic.NewRangeQuery("end_time").Gt(time.Now().Unix()))

	res, err := esClient.Search(utils.UnMarshal(models.ApiOrder{})).
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
		var apiOrder models.ApiOrder
		data, _ := item.Source.MarshalJSON()
		json.Unmarshal(data, &apiOrder)
		apiOrder.Id = item.Id
		r = append(r, apiOrder)
	}
	return
}
