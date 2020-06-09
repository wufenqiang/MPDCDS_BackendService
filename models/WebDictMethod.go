package models

//数据获取方式码表
type WebDictMethod struct {
	Id   string `json:"-"`    //系统ID
	Name string `json:"name"` //1,FTP	2,HTTP	3,其它方式
}
