package models

import "time"

//菜单表
type WebMenu struct {
	Id         string    `json:"-"`           //系统ID
	MenuName   string    `json:"menu_name"`   //菜单名称
	AccessUrl  string    `json:"access_url"`  //访问路径
	Level      int       `json:"level"`       //级别
	SortNum    int       `json:"sort_num"`    //排序
	Icon       string    `json:"icon"`        //图标
	CreateTime time.Time `json:"create_time"` //创建时间
	Remark     string    `json:"Remark"`      //备注
}
