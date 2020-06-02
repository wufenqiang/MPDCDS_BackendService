package models

import (
	"github.com/jinzhu/gorm"
)

type Method_dict struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Name int `gorm:"type:varchar(10)"`
}
