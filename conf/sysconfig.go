package conf

import (
	"github.com/json-iterator/go"
	"io/ioutil"
)

var Sysconfig = &sysconfig{}

func init() {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Sys config read err")
	}
	err = jsoniter.Unmarshal(b, Sysconfig)
	if err != nil {
		panic(err)
	}
}

type sysconfig struct {
	//端口
	Port string `json:"Port"`

	//  "DBIp": "220.243.129.248",
	//  "DBName": "iris_gorm_dev",
	//  "DBUserName": "iris_gorm",
	//  "DBPassword": "123456",
	//  "DBPort": "3306",
	//mysql信息
	//DBUserName string `json:"DBUserName"`
	//DBPassword string `json:"DBPassword"`
	//DBIp       string `json:"DBIp"`
	//DBPort     string `json:"DBPort"`
	//DBName     string `json:"DBName"`

	//日志存储地址
	LoggerPath  string `json:"LoggerPath"`
	LoggerLevel string `json:"LoggerLevel"`

	//thrift 服务ip:port
	NetworkAddr string `json:"NetworkAddr"`

	ESURL string `json:"ESURL"`
}
