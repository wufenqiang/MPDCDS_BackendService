package models

import "time"

//接入数据信息表
type ApiDataInfo struct {
	Id             string    `json:"id"`              //系统ID
	DataName       string    `json:"data_name"`       //数据类型名称
	DataCode       string    `json:"data_code"`       //数据编码
	ServerAddress  string    `json:"server_address"`  //服务器地址
	StorageAddress string    `json:"storage_address"` //数据存储地址
	Name           string    `json:"name"`            //文件名称、表名
	Frequency      string    `json:"frequency"`       //更新频率
	Source         string    `json:"source"`          //数据来源（提供方）
	FileSize       int64     `json:"file_size"`       //文件大小（字节）
	Reliability    string    `json:"reliability"`     //数据可靠性（0已签约、1未签约）
	CreateTime     time.Time `json:"create_time"`     //创建时间
	Remark         string    `json:"remark"`          //备注
	DataType       string    `json:"data_type"`       //数据类型（0文件、1数据）
	Feature        string    `json:"feature"`         //要素信息
	IsOpen         string    `json:"is_open"`         //是否开放（对外提供 0是、1否  用户是否可见）
	ValidTime      int       `json:"value_time"`      //数据保留时间（天）
	Admode         string    `json:"admode"`          //采集方式
	LinkMan        string    `json:"link_man"`        //联系人
	ConvergeStatus string    `json:"converge_status"` //内部数据汇聚状态（1是、2否  给数据采集客户端使用）
	AccessState    string    `json:"access_state"`    //上游数据接入状态（1未接入、2试运行、3已稳定   判断上游接入数据是否稳定）
	DataLevel      string    `json:"DataLevel"`       //数据级别（1 2 3   3级不需要审批）
	FileIndexName  string    `json:"file_index_name"` //文件索引名称
}
