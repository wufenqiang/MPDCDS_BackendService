package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Hf_platform_rolemenu struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Roleid     int       `gorm:"type:int(11)"`
	Menuid     int       `gorm:"type:int(11)"`
	Createtime time.Time `gorm:"type:datetime(0)"`
}
