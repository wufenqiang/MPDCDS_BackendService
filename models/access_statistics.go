package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Access_statistics struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Userid     int       `gorm:"type:int(11)"`
	Accesstime time.Time `gorm:"type:datetime(0)"`
	Accessid   int       `gorm:"type:int(11)"`
	Fileid     string    `gorm:"type:int(11)"`
}
