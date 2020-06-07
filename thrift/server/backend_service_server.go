package server

import (
	"MPDCDS_BackendService/conf"
	esdatasource "MPDCDS_BackendService/datasource/elasticsearch"
	"MPDCDS_BackendService/logger"
	"MPDCDS_BackendService/models"
	"MPDCDS_BackendService/thrift/MPDCDS_BackendService"
	"MPDCDS_BackendService/utils"
	"context"
	"encoding/json"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type MPDCDS_BackendServiceImpl struct {
}

func (this *MPDCDS_BackendServiceImpl) Auth(ctx context.Context, user string, password string) (r *MPDCDS_BackendService.Auth, err error) {
	//todo 查询es验证用户名和密码是否合法
	r = MPDCDS_BackendService.NewAuth()
	esclient := esdatasource.GetESClient()
	u := elastic.NewQueryStringQuery(user)     //"username:hfcmjt"
	p := elastic.NewQueryStringQuery(password) //"password:123456789"
	res, err := esclient.Search("web_user").Query(u).Query(p).Do(context.Background())
	if err != nil {
		r.Status = -1
		r.Msg = "根据用户名和密码查询用户信息失败！"
		logger.GetLogger().Error("根据用户名和密码查询hf_platform_user失败，异常信息" + err.Error())
		return
	}
	if res.Hits.TotalHits.Value > 0 {
		var t models.WebUser
		for _, hit := range res.Hits.Hits {
			err := json.Unmarshal(hit.Source, &t) //另外一种取数据的方法
			if err != nil {
				r.Status = -1
				r.Msg = "数据格式化失败！"
				logger.GetLogger().Error("数据格式化失败！")
			}
			t.Id = hit.Id
			logger.GetLogger().Info("Hf_platform_user", zap.String("Id", hit.Id))
			break
		}
		//用户名和userId合法,使用jwt生成token,userId为es中查询的用户信息主键
		token, err := utils.GenerateToken(user, t.Id)
		if err != nil {
			r.Status = -1
			r.Msg = "生成token失败！"
			logger.GetLogger().Error("生成token失败！")
		} else {
			r.Status = 0
			r.Token = token
			r.Msg = "用户名和密码正确！"
		}
	} else {
		r.Status = -1
		r.Msg = "用户名或密码错误！"
	}
	return
}

func (this *MPDCDS_BackendServiceImpl) Lists(ctx context.Context, token string, pwd string) (r *MPDCDS_BackendService.FileDirInfo, err error) {
	//验证token是否有效
	m := make(map[string]string)
	isValid, err := utils.VerifyToken(m, token)
	//合法
	if isValid {
		//todo 根据用户信息、当前目录从es中查询数据目录和数据列表
		logger.GetLogger().Info("用户id:" + m["id"] + "=========用户名称" + m["username"])
	}
	return
}

func (this *MPDCDS_BackendServiceImpl) File(ctx context.Context, pwd string, path string) (r *MPDCDS_BackendService.FileInfo, err error) {
	//todo 根据当前目录和文件名称查询文件真实地址

	return
}

//启动thrift server服务
func InitMpdcdsBackendServiceServer() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(conf.Sysconfig.NetworkAddr)
	logger.GetLogger().Info("thrift server start.......")
	if err != nil {
		logger.GetLogger().Error(err.Error())
	}
	handler := &MPDCDS_BackendServiceImpl{}
	processor := MPDCDS_BackendService.NewMPDCDS_BackendServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	logger.GetLogger().Info("thrift server in" + conf.Sysconfig.NetworkAddr)
	server.Serve()
}
