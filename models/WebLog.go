package models

import "time"

//日志记录表
type WebLog struct {
	//Id         string    `json:"id"`          //系统ID
	UserName   string    `json:"user_name"`   //用户名称
	Operate    string    `json:"operate"`     //操作（查询、编辑、新增、删除）
	CataLog    string    `json:"cata_log"`    //对象
	CreateTime time.Time `json:"create_time"` //创建时间
	UserId     string    `json:"user_id"`     //用户表
}
