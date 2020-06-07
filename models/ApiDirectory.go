package models

import "time"

//目录信息表
type ApiDirectory struct {
	DirId      string    `json:"dir_id"`      //数据目录ID
	FatherId   string    `json:"father_id"`   //父目录ID
	CreateTime time.Time `json:"create_time"` //创建时间
	//todo 数据ID的含义是？
	DataId string `json:"data_id"` //数据ID
	Path   string `json:"path"`    //全路径
}
