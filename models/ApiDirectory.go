package models

import "time"

//目录信息表
type ApiDirectory struct {
	Id            string    `json:"id"`              //系统ID
	ParentDir     string    `json:"parent_dir"`      //上级目录地址(/ocf)
	CurrentDir    string    `json:"current_dir"`     //当前目录地址(/ocf/1h)
	CreateTime    time.Time `json:"create_time"`     //创建时间
	FileIndexName string    `json:"file_index_name"` //文件索引名称(api_file_类型)
}
