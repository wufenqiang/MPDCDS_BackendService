package server

import (
	"MPDCDS_BackendService/conf"
	esdatasource "MPDCDS_BackendService/datasource/elasticsearch"
	"MPDCDS_BackendService/logger"
	"MPDCDS_BackendService/models"
	"MPDCDS_BackendService/service"
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
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewTermQuery("user_name.keyword", user), elastic.NewTermsQuery("password.keyword", password))
	res, err := esclient.Search(utils.UnMarshal(models.WebUser{})).Query(boolQuery).Do(context.Background())
	if err != nil {
		r.Status = -1
		r.Msg = "根据用户名和密码查询用户信息失败！"
		logger.GetLogger().Error("根据用户名和密码查询web_user失败，异常信息" + err.Error())
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
			logger.GetLogger().Info("web_user", zap.String("Id", hit.Id))
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
	r = MPDCDS_BackendService.NewFileDirInfo()
	//验证token是否有效
	m := make(map[string]string)
	isValid, err := utils.VerifyToken(m, token)

	//合法
	userIdLog := zap.String("userId", m["id"])
	userNameLog := zap.String("username", m["username"])
	accessPathLog := zap.String("accessPath", pwd)
	if isValid {
		logger.GetLogger().Info("user access", userIdLog, userNameLog, accessPathLog)
	} else {
		r.Status = -1
		r.Msg = "User authentication failed"
		logger.GetLogger().Error("User authentication failed", userIdLog, userNameLog, accessPathLog)
		return
	}

	fileService := service.NewApiFileService()
	r.Status = 0
	r.Data = fileService.GetFileByPath(m["id"], pwd)
	r.Msg = "Get list information succeeded"
	return
}

func (this *MPDCDS_BackendServiceImpl) DirAuth(ctx context.Context, token string, abspath string) (r *MPDCDS_BackendService.DirAuth, err error) {
	//todo 判断当前用户是否有权限访问该目录
	r = MPDCDS_BackendService.NewDirAuth()
	//验证token是否有效
	m := make(map[string]string)
	isValid, err := utils.VerifyToken(m, token)

	//合法
	userIdLog := zap.String("userId", m["id"])
	userNameLog := zap.String("username", m["username"])
	accessPathLog := zap.String("accessPath", abspath)
	if isValid {
		logger.GetLogger().Info("user access", userIdLog, userNameLog, accessPathLog)
	} else {
		r.Status = -1
		r.Msg = "User authentication failed"
		logger.GetLogger().Error("User authentication failed", userIdLog, userNameLog, accessPathLog)
		return
	}

	apiFileService := service.NewApiFileService()
	status, msg := apiFileService.ValidDirByUserOrder(m["id"], abspath)

	r.Status = status
	r.Msg = msg
	return
}

//启动thrift server服务
func InitMpdcdsBackendServiceServer() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(conf.Sysconfig.ThriftHost + ":" + conf.Sysconfig.ThriftPort)

	logger.GetLogger().Info("thrift server start.......")
	if err != nil {
		logger.GetLogger().Error(err.Error())
	}
	handler := &MPDCDS_BackendServiceImpl{}
	processor := MPDCDS_BackendService.NewMPDCDS_BackendServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	serverTransport.Addr().String()
	logger.GetLogger().Info("thrift server in " + conf.Sysconfig.ThriftHost + ":" + conf.Sysconfig.ThriftPort)
	server.Serve()
}
