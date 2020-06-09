package models

import "time"

//用户角色关系表
type WebUserRole struct {
	Id         string    `json:"id"`          //系统ID
	UserID     string    `json:"user_id"`     //用户表主键
	RoleId     string    `json:"role_id"`     //角色表主键
	CreateTime time.Time `json:"create_time"` //创建时间
}
