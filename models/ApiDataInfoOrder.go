package models

//接入数据信息订单信息关系表
type ApiDataInfoOrder struct {
	Id       string `json:"id"`        //系统ID
	OrderId  string `json:"order_id"`  //订单表主键
	AccessId string `json:"access_id"` //接入数据信息表主键
}
