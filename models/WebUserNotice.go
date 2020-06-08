package models

import "time"

//用户通知信息表关系表
type WebUserNotice struct {
	//Id         string    `json:"id"`          //系统ID
	UserId     string    `json:"user_id"`     //用户表主键
	NoticeId   string    `json:"notice_id"`   //通知信息表主键
	CreateTime time.Time `json:"create_time"` //创建时间
}
