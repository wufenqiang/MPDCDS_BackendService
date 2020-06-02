package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Orderinfo struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Orderno        string    `gorm:"type:varchar(20)"`
	Method_id      int       `gorm:"type:int(11)"`
	Createtime     time.Time `gorm:"type:datetime(0)"`
	Starttime      time.Time `gorm:"type:datetime(0)"`
	Endtime        time.Time `gorm:"type:datetime(0)"`
	Server_address string    `gorm:"type:varchar(50)"`
	Name           string    `gorm:"type:varchar(100)"`
	Descr          string    `gorm:"type:varchar(500)"`
	User           string    `gorm:"type:varchar(100)"`
	Remark         string    `gorm:"type:varchar(500)"`
	Status         string    `gorm:"type:char(1)"`
	Duration       int       `gorm:"type:int(11)"`
	Instanceid     int       `gorm:"type:int(11)"`
}
