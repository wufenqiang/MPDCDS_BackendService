package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Download struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Accessid     int       `gorm:"type:int(11)"`
	Fileid       int       `gorm:"type:int(11)"`
	Createtime   time.Time `gorm:"type:datetime(0)"`
	Downloadtime time.Time `gorm:"type:datetime(0)"`
	Userid       int       `gorm:"type:int(11)"`
}
