package models

import "time"

//文件下载信息记录表
type ApiDownload struct {
	Id         string    `json:"-"`           //系统ID
	AccessId   string    `json:"access_id"`   //接入数据信息表主键
	FileId     string    `json:"file_id"`     //文件信息表主键
	StartTime  time.Time `json:"start_time"`  //文件开始下载时间时间
	EndTime    time.Time `json:"end_time"`    //更新时间
	UserId     string    `json:"user_id"`     //用户信息表主键
	CreateTime time.Time `json:"create_time"` //创建时间
}
