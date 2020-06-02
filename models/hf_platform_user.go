package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Hf_platform_user struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Username   string    `gorm:"type:varchar(20);not null;"`
	Password   string    `gorm:"type:varchar(20);not null;"`
	Createtime time.Time `gorm:"type:datetime(0)"`
	Realname   string    `gorm:"type:varchar(20)"`
	Phone      string    `gorm:"type:varchar(11)"`
	Email      string    `gorm:"type:varchar(20)"`
	Appkey     string    `gorm:"type:varchar(36)"`
	Status     string    `gorm:"type:char(1)"`
	Reamark    string    `gorm:"type:varchar(300)"`
}
