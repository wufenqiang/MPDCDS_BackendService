package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Access_data_info struct {
	gorm.Model
	//id   string `gorm:"type:int(11);not null;"`
	Dataname        string    `gorm:"type:varchar(36)"`
	Datacode        string    `gorm:"type:varchar(36)"`
	Server_address  string    `gorm:"type:varchar(50)"`
	Storage_address string    `gorm:"type:varchar(100)"`
	Name            string    `gorm:"type:varchar(50)"`
	Frequency       string    `gorm:"type:varchar(50)"`
	Source          string    `gorm:"type:varchar(50)"`
	Filesize        string    `gorm:"type:bigint(11)"`
	Reliability     string    `gorm:"type:char(50)"`
	Createtime      time.Time `gorm:"type:datetime(0)"`
	Remark          string    `gorm:"type:varchar(300)"`
	Datatype        string    `gorm:"type:char(1)"`
	Feature         string    `gorm:"type:text(0)"`
	Isopen          string    `gorm:"type:char(1)"`
	Validtime       int       `gorm:"type:int(11)"`
	Admode          int       `gorm:"type:varchar(30)"`
	Linkman         int       `gorm:"type:varchar(20)"`
	Convergestatus  int       `gorm:"type:char(1)"`
	Accessstate     int       `gorm:"type:char(1)"`
	Datalevel       int       `gorm:"type:char(1)"`
}
