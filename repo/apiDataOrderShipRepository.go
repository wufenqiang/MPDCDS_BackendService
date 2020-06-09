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

type ApiDataOrderShipRepository interface {
	//根据用户ID获取订单
	GetDataOrderShipListByOrderId(orderIds []interface{}) []string
}

func NewApiDataOrderShipRepository() ApiDataOrderShipRepository {
	return &apiDataOrderShipRepository{}
}

type apiDataOrderShipRepository struct{}

func (a apiDataOrderShipRepository) GetDataOrderShipListByOrderId(orderIds []interface{}) (r []string) {
	esClient := esdatasource.GetESClient()

	if orderIds == nil {
		return
	}

	boolQ := elastic.NewBoolQuery()
	if len(orderIds) > 0 {
		boolQ.Must(elastic.NewTermsQuery("order_id.keyword", orderIds...))
	}

	res, err := esClient.Search(utils.UnMarshal(models.ApiDataInfoOrder{})).
		Size(10000).
		From(0).
		Query(boolQ).Do(context.Background())

	if err != nil {
		logger.GetLogger().Error("GetDataOrderShipListByOrderId", zap.Error(err))
	}

	if res == nil {
		return
	}

	for _, item := range res.Hits.Hits {
		var apiDataInfoOrder models.ApiDataInfoOrder
		data, _ := item.Source.MarshalJSON()
		json.Unmarshal(data, &apiDataInfoOrder)
		r = append(r, apiDataInfoOrder.AccessId)
	}
	return
}
