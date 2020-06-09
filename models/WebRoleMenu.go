package models

import "time"

//角色菜单表
type WebRoleMenu struct {
	Id         string    `json:"id"`          //系统ID
	RoleId     string    `json:"role_id"`     //角色表主键
	MenuId     string    `json:"menu_id"`     //菜单表主键
	CreateTime time.Time `json:"create_time"` //创建时间
}
