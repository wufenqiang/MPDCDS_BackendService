package models

import "time"

//订单信息表
type ApiOrder struct {
	Id            string    `json:"id"`             //系统ID
	OrderNum      string    `json:"order_num"`      //订单号
	MethodId      string    `json:"method_id"`      //数据获取方式编码
	CreateTime    time.Time `json:"create_time"`    //创建时间
	StartTime     time.Time `json:"start_time"`     //开始使用时间
	EndTime       time.Time `json:"end_time"`       //结束时间
	ServerAddress string    `json:"server_address"` //下游获取服务器地址（是否需要限制该服务器访问）
	Name          string    `json:"name"`           //业务名称
	Descr         string    `json:"descr"`          //业务描述
	User          string    `json:"user"`           //服务对象
	Remark        string    `json:"remark"`         //备注
	Status        string    `json:"status"`         //审批状态（1待审批、2审批中、3审批完成、4审批不通过）
	Duration      int       `json:"duration"`       //数据时长
	InstanceId    string    `json:"instance_id"`    //流程ID
	AccessId      string    `json:"access_id"`      //接入数据基本信息表主键
}
