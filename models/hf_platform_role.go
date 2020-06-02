package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Hf_platform_role struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Rolename   string    `gorm:"type:varchar(20)"`
	Rolecode   string    `gorm:"type:varchar(36)"`
	Createtime time.Time `gorm:"type:datetime(0)"`
	Remark     string    `gorm:"type:varchar(300)"`
}
