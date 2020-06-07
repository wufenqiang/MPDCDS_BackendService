package models

import (
	"time"
)

type Hf_platform_user struct {
	Id         string    `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Createtime time.Time `json:"createtime"`
	Realname   string    `json:"realname"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Appkey     string    `json:"appkey"`
	Status     string    `json:"status"`
	Reamark    string    `json:"reamark"`
}
