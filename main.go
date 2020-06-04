package main

import (
	"MPDCDS_BackendService/conf"
	thelogger "MPDCDS_BackendService/logger"
	"MPDCDS_BackendService/route"
	"flag"
	"github.com/kataras/iris"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	flag.Parse()
	app := newApp()
	route.InitRouter(app)
	//初始化日志
	logger = thelogger.InitLog(conf.Sysconfig.LoggerPath, conf.Sysconfig.LoggerLevel)
	logger.Info("start print logger......")

	err := app.Run(iris.Addr(":"+conf.Sysconfig.Port), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
	logger.Info("end print logger......")
}

//noinspection GoTypesCompatibility
func newApp() *iris.Application {
	app := iris.New()
	//app.Configure(iris.WithOptimizations)
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowCredentials: true,
	//	AllowedHeaders:   []string{"*"},
	//})
	//app.Use(crs)
	app.AllowMethods(iris.MethodOptions)
	return app
}
