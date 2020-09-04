package main

import (
	"MPDCDS_BackendService/src/conf"
	"MPDCDS_BackendService/src/logger"
	"MPDCDS_BackendService/src/thrift/thrift-server"
	"flag"
)

func main() {
	flag.Parse()
	//app := newApp()
	//route.InitRouter(app)
	//初始化日志
	logger.GetLogger().Info(conf.ProjectName + " start ...")
	//启动 thrift server
	thrift_server.InitThriftServer()
}

//noinspection GoTypesCompatibility
//func newApp() *iris.Application {
//	app := iris.New()
//	//app.Configure(iris.WithOptimizations)
//	//crs := cors.New(cors.Options{
//	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
//	//	AllowCredentials: true,
//	//	AllowedHeaders:   []string{"*"},
//	//})
//	//app.Use(crs)
//	app.AllowMethods(iris.MethodOptions)
//	return app
//}
