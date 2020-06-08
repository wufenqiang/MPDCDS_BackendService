package models

import "time"

//角色表
type WebRole struct {
	//Id         string    `json:"id"`          //系统ID
	RoleName   string    `json:"role_name"`   //角色名称
	RoleCode   string    `json:"role_code"`   //角色编码
	CreateTime time.Time `json:"create_time"` //创建时间
	Remark     string    `json:"remark"`      //备注
}
