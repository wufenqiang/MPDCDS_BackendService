package models

import "time"

//数据访问统计表
//todo 与用户下载记录表重复？
type WebAccessStatistics struct {
	Id         string    `json:"id"`          //系统ID
	UserId     string    `json:"user_id"`     //用户表主键
	AccessTime time.Time `json:"access_time"` //访问时间
	AccessId   string    `json:"access_id"`   //接入数据信息主键
	FileId     string    `json:"FileId"`      //文件信息表主键
}
