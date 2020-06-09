package models

import "time"

//通知信息表
type WebNotice struct {
	Id         string    `json:"id"`          //系统ID
	Title      string    `json:"title"`       //标题
	Details    string    `json:"details"`     //内容
	UserName   string    `json:"user_name"`   //发送人
	CreateTime time.Time `json:"create_time"` //创建时间
	UserId     string    `json:"user_id"`     //用户表主键
}
