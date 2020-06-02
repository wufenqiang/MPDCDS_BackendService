package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Hf_platform_usernotice struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Userid     int       `gorm:"type:int(11)"`
	Noticeid   int       `gorm:"type:int(11)"`
	Createtime time.Time `gorm:"type:datetime(0)"`
}
