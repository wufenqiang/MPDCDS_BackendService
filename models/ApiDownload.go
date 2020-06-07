package models

import "time"

//用户下载记录信息表
type ApiDownload struct {
	Id         string    `json:"id"`          //系统ID
	AccessId   string    `json:"access_id"`   //接入数据信息表主键
	FileId     string    `json:"file_id"`     //文件信息表主键
	CreateTime time.Time `json:"create_time"` //创建时间
	//todo 下载时间与创建时间不同？
	DownloadTime time.Time `json:"download_time"` //下载时间
	UserId       string    `json:"user_id"`       //用户信息表主键
}
