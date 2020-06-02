package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Hf_platform_memu struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Meauname   string    `gorm:"type:varchar(50)"`
	Accessurl  string    `gorm:"type:varchar(100)"`
	Level      int       `gorm:"type:int(11)"`
	Sortnum    int       `gorm:"type:int(11)"`
	Icon       string    `gorm:"type:varchar(50)"`
	Createtime time.Time `gorm:"type:datetime(0)"`
	Remark     string    `gorm:"type:varchar(300)"`
}
