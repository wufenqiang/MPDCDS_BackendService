package models

import "time"

//文件信息表
type ApiFile struct {
	Id         string    `json:"id"`          //系统ID
	FileName   string    `json:"file_name"`   //文件名称
	AccessId   string    `json:"access_id"`   //接入数据信息表主键
	FileSize   int64     `json:"file_size"`   //文件大小（字节）
	CreateTime time.Time `json:"create_time"` //创建时间
	DirId      string    `json:"dir_id"`      //目录Id
}
