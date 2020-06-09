package models

import "time"

//接入基本数据信息扩展表
type ApiDataInfoExt struct {
	Id         string    `json:"id"`          //系统ID
	FieldName  string    `json:"field_name"`  //字段名称
	FieldValue string    `json:"field_value"` //字段值
	FieldType  string    `json:"field_type"`  //字段类型
	FieldDesc  string    `json:"field_desc"`  //字段描述
	CreateTime time.Time `json:"create_time"` //创建时间
	SortNum    int       `json:"sort_num"`    //排序展示值
	GroupName  string    `json:"group_name"`  //分组名称
	AccessId   string    `json:"access_id"`   //接入信息表主键
}
