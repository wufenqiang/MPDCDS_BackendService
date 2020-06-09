package models

import "time"

//用户表
type WebUser struct {
	Id         string    `json:"-"`           //系统ID
	UserName   string    `json:"user_name"`   //用户名称
	Password   string    `json:"password"`    //密码
	CreateTime time.Time `json:"create_time"` //创建时间
	RealName   string    `json:"real_name"`   //真实姓名
	Phone      string    `json:"phone"`       //手机号
	Email      string    `json:"email"`       //邮箱
	AppKey     string    `json:"app_key"`     //唯一标识
	Status     string    `json:"status"`      //0启用、1停用
	Remark     string    `json:"remark"`      //备注
}
