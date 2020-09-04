package repo

import (
	"MPDCDS_BackendService/src/datasource"
	"MPDCDS_BackendService/src/logger"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"
)

var ApiOrderRepo = new(apiOrderRepo)

type apiOrderRepo struct{}

//根据用户ID获取订单
func (a *apiOrderRepo) GetOrderByUserId(userId string) (r []models.Api_order) {
	if userId == "" {
		return
	}

	dataClient := datasource.MySQL.GetDB()

	e := dataClient.
		Find(&r, "user_id = ?", userId).
		Error

	if e != nil {
		logger.GetLogger().Error(e.Error())
	}

	return
}

//根据订单ID获取数据
func (a *apiOrderRepo) GetDataCoreByOrderId(orderIds []interface{}) (datacodes []string) {

	if orderIds == nil || len(orderIds) < 1 {
		return
	}

	dataClient := datasource.MySQL.GetDB()
	for _, orderid := range orderIds {
		var apiorderplans []models.Api_order_plan
		e := dataClient.Find(&apiorderplans, "order_id = ?", orderid).Error
		if e == nil {
			for _, apiorderplan := range apiorderplans {
				datacodes = append(datacodes, apiorderplan.DataCode)
			}
		}
	}

	return
}

func (a *apiOrderRepo) GetDataCoreByOrderNo(orderno interface{}) (datacodes []string) {
	dataClient := datasource.MySQL.GetDB()

	var apiorder models.Api_order
	dataClient.Find(&apiorder, "order_no = ?", orderno)
	orderid := apiorder.OrderId
	var apiorderplans []models.Api_order_plan
	dataClient.Find(&apiorderplans, "order_id =?", orderid)

	for _, apiorderplan := range apiorderplans {
		datacodes = append(datacodes, apiorderplan.DataCode)
	}

	return
}

func (a *apiOrderRepo) ValidOrderByUserId(userId string, orderNo string) bool {
	if userId == "" {
		return false
	}

	dataClient := datasource.MySQL.GetDB()
	var r models.Api_order

	err := dataClient.
		Find(&r, "user_id = "+userId+" and order_no= ?", orderNo).Error

	if err == nil {
		return true
	} else {
		return false
	}
}
func (a *apiOrderRepo) ValidDataByUserIdOrderNo(userId string, orderNo string, datacode string) bool {
	if userId == "" {
		return false
	}

	dataClient := datasource.MySQL.GetDB()
	var r models.Api_order

	err0 := dataClient.
		Find(&r, "user_id = "+userId+" and order_no= ?", orderNo).Error

	if err0 == nil {
		var apiorderplans []models.Api_order_plan
		err1 := dataClient.Find(&apiorderplans, "order_id =?", r.OrderId).Error

		if err1 == nil {
			return true
		} else {
			return false
		}

	} else {
		return false
	}
}

//根据数据Id获取数据
//func (a *apiOrderRepo) GetDataCodeByDataId(ids []string) (r []bak.ApiDataInfo) {
//	dataClient := datasource.ES.GetDB()
//
//	if ids == nil || len(ids) < 1 {
//		return
//	}
//
//	boolQ := elastic.NewBoolQuery()
//	boolQ.Filter(elastic.NewIdsQuery().Ids(ids...))
//
//	res, err := dataClient.Search(proutils.UnMarshal(bak.ApiDataInfo{})).
//		Size(10000).
//		From(0).
//		Query(boolQ).Do(context.Background())
//
//	if err != nil {
//		logger.GetLogger().Error("GetApiDataInfoById", zap.Error(err))
//		return
//	}
//
//	if res == nil {
//		return
//	}
//
//	for _, item := range res.Hits.Hits {
//		var apiApiDataInfo bak.ApiDataInfo
//		data, _ := item.Source.MarshalJSON()
//		json.Unmarshal(data, &apiApiDataInfo)
//		apiApiDataInfo.Id = item.Id
//		r = append(r, apiApiDataInfo)
//	}
//	return
//}
