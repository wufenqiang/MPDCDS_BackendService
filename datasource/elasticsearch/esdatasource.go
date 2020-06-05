package esdatasource

import (
	"MPDCDS_BackendService/conf"
	"MPDCDS_BackendService/logger"
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"os"
)

var esclient *elastic.Client

func GetESClient() *elastic.Client {
	return esclient
}

func init() {
	esurl := conf.Sysconfig.ESURL
	//esurl := "http://elastic:weather.com.cn@220.243.130.220:9200"

	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	esclient, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(esurl))
	if err != nil {
		panic(err)
	}
	info, code, err := esclient.Ping(esurl).Do(context.Background())
	if err != nil {
		panic(err)
	}

	//fmt.Println(info,code)
	logger.GetLogger().Info("Elasticsearch connect ", zap.Int("code", code), zap.String("Version", info.Version.Number))

	/*	esversion, err := esclient.ElasticsearchVersion(conf.Sysconfig.ESURL)
		if err != nil {
			panic(err)
		}
		logger.GetLogger().Info("Elasticsearch version ", zap.String("version",esversion))*/
}
