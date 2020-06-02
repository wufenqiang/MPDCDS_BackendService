package models

import "github.com/jinzhu/gorm"

type Access_data_info_order struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Orderid  int `gorm:"type:int(11)"`
	Accessid int `gorm:"type:int(11)"`
}
