package datasource

import (
	"MPDCDS_BackendService/src/conf"
	"MPDCDS_BackendService/src/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

var MySQL = new(mysql0)

type mysql0 struct {
}

var mysqlclient *gorm.DB

func (*mysql0) GetDB() *gorm.DB {
	return mysqlclient
}

func init() {
	path := strings.Join([]string{conf.Sysconfig.MySQLUserName, ":", conf.Sysconfig.MySQLPassWord, "@(", conf.Sysconfig.MySQLIp, ":", conf.Sysconfig.MySQLPort, ")/", conf.Sysconfig.MySQLDBName, "?charset=utf8&parseTime=true"}, "")
	var err error
	mysqlclient, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	mysqlclient.SingularTable(true)
	mysqlclient.DB().SetConnMaxLifetime(1 * time.Second)
	mysqlclient.DB().SetMaxIdleConns(20)   //最大打开的连接数
	mysqlclient.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	mysqlclient.SingularTable(true)        //表生成结尾不带s
	// 启用Logger，显示详细日志
	mysqlclient.LogMode(true)
	logger.GetLogger().Info("MySQL初始化完成！")
	//Createtable()
}
