package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Hf_platform_notice struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Title      string    `gorm:"type:varchar(30)"`
	Detail     string    `gorm:"type:varchar(500)"`
	Username   string    `gorm:"type:varchar(20)"`
	Createtime time.Time `gorm:"type:datetime(0)"`
}
