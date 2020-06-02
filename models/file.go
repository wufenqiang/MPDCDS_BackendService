package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type File struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	FileName   string    `gorm:"type:varchar(100)"`
	Accessid   int       `gorm:"type:int(11)"`
	Filesize   int       `gorm:"type:bigint(11)"`
	Createtime time.Time `gorm:"type:datetime(0)"`
}
