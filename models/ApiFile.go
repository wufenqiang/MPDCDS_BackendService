package models

import "time"

//文件信息表
type ApiFile struct {
	Id          string    `json:"-"`            //系统ID
	FileName    string    `json:"file_name"`    //文件名称
	AccessId    string    `json:"access_id"`    //接入数据信息表主键
	FileSize    int64     `json:"file_size"`    //文件大小（字节）
	CreateTime  time.Time `json:"create_time"`  //创建时间
	DirId       string    `json:"dir_id"`       //根据最小目录地址生成唯一标识
	FileAddress string    `json:"file_address"` //文件地址  本地地址（file:///） http://ip:port/xxxx
	CollectTime time.Time `json:"collect_time"` //采集时间
	UpdateTime  time.Time `json:"updatetime"`   //更新时间
	Status      string    `json:"status"`       //0可用
}
