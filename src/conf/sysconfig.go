package conf

import (
	"fmt"
	"github.com/json-iterator/go"
	proutils "gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/utils"
	"io/ioutil"
	"time"
)

type sysconfig struct {
	//ProjectName string `json:"ProjectName"`

	//端口
	HttpsPort string `json:"HttpsPort"`

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
	LoggerLevel string `json:"LoggerLevel"`

	//thrift 服务ip:port
	ThriftHost string `json:"ThriftHost"`
	ThriftPort string `json:"ThriftPort"`

	ESURL string `json:"ESURL"`

	MySQLIp       string `json:"MySQLIp""`
	MySQLPort     string `json::"MySQLPort"`
	MySQLDBName   string `json:"MySQLDBName"`
	MySQLUserName string `json:"MySQLUserName"`
	MySQLPassWord string `json:"MySQLPassWord"`
}

var Sysconfig = &sysconfig{}

const ProjectName = "MPDCDS_BackendService"
const Layout = "2006-01-02 15:04:05" //时间格式
const TheLocation = "Asia/Shanghai"

var Loc, _ = time.LoadLocation(TheLocation)
var TimeStamp = time.Now()

func init() {
	fmt.Println(ProjectName)
	dir := proutils.ProjectLocation(ProjectName)
	conffile := dir + "/config.json"
	ReadConf(conffile, Sysconfig)
}
func LocalProjectPath() string {
	return proutils.ProjectLocation(ProjectName)
}

func ReadConf(conffile string, thesysconf interface{}) {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile(conffile)
	if err != nil {
		panic(conffile + "Sys config read err")
	}
	err = jsoniter.Unmarshal(b, thesysconf)
	if err != nil {
		panic(err)
	}
}
