package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Hf_platform_log struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Username   string    `gorm:"type:varchar(20)"`
	Operate    string    `gorm:"type:varchar(20)"`
	Catalog    string    `gorm:"type:varchar(50)"`
	Createtime time.Time `gorm:"type:datetime(0)"`
}
