package thrift_server

import (
	"MPDCDS_BackendService/src/conf"
	"MPDCDS_BackendService/src/logger"
	"MPDCDS_BackendService/src/repo"
	"MPDCDS_BackendService/src/service"
	"MPDCDS_BackendService/src/utils"
	"github.com/apache/thrift/lib/go/thrift"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/thrift/thriftcore"
	"strconv"

	"context"
	"go.uber.org/zap"
)

type MPDCDSProServiceImpl struct {
}

//已修改
//todo 查询es验证用户名和密码是否合法
func (this *MPDCDSProServiceImpl) Auth(ctx context.Context, authInfo *thriftcore.AuthInfo) (r *thriftcore.AuthReturn, err error) {
	r = thriftcore.NewAuthReturn()

	user := authInfo.User
	password := authInfo.Password

	apiuser, e := repo.
		ApiUserRepo.
		GetUserByUserName(user)

	//var apiuser models.Api_user
	//e := mysqlclient.Find(&apiuser, "username = ?", user).Error
	if e == nil {
		if apiuser.Password == password {
			//用户名和userId合法,使用jwt生成token,userId为es中查询的用户信息主键
			token, err := utils.GenerateToken(user, strconv.Itoa(apiuser.Id))
			if err != nil {
				r.Status = 2
				r.Msg = thriftcore.AuthReturnCodeMap[r.Status]
				logger.GetLogger().Error(r.Msg)
			} else {
				r.Status = 0
				r.Token = token
				r.Msg = thriftcore.AuthReturnCodeMap[r.Status]
			}
		} else {
			r.Status = 3
			r.Msg = thriftcore.AuthReturnCodeMap[r.Status]
		}
	} else {
		r.Status = 1
		r.Msg = thriftcore.AuthReturnCodeMap[r.Status]
	}

	return r, e
}

//todo 目录信息和列表信息接口
func (this *MPDCDSProServiceImpl) Lists(ctx context.Context, listInfo *thriftcore.ListsInfo) (r *thriftcore.ListsReturn, err error) {
	token := listInfo.Token
	pwd := listInfo.Pwd

	r = thriftcore.NewListsReturn()
	//验证token是否有效
	verifyTokenReturn := make(map[string]string)
	isValid, err := utils.VerifyToken(verifyTokenReturn, token)
	userId := verifyTokenReturn["id"]
	username := verifyTokenReturn["username"]

	//合法
	if isValid {
		fileService := service.NewApiFileService()
		r.Status = 0
		r.Data = fileService.GetFileByPath(userId, pwd)
		r.Msg = thriftcore.ListsReturnCodeMap[r.Status]
		logger.GetLogger().Info("Failed! User[" + userId + "][" + username + "][" + pwd + "]")
		return
	} else {
		r.Status = 1
		r.Msg = thriftcore.ListsReturnCodeMap[r.Status]
		logger.GetLogger().Error("Successed! User[" + userId + "][" + username + "][" + pwd + "]")
		return
	}

}

//todo 判断当前用户是否有权限访问该目录
func (this *MPDCDSProServiceImpl) DirAuth(ctx context.Context, dirAuthInfo *thriftcore.DirAuthInfo) (dirAuthReturn *thriftcore.DirAuthReturn, err error) {
	token := dirAuthInfo.Token
	absPath := dirAuthInfo.AbsPath

	dirAuthReturn = thriftcore.NewDirAuthReturn()
	//验证token是否有效
	m := make(map[string]string)
	isValid, err := utils.VerifyToken(m, token)

	userId := m["id"]
	username := m["username"]
	//合法
	if isValid {
		logger.GetLogger().Info("Token Successed. user[" + userId + "][" + username + "][" + absPath + "]")
	} else {
		dirAuthReturn.Status = 1
		dirAuthReturn.Msg = thriftcore.DirAuthReturnCodeMap[dirAuthReturn.Status]
		logger.GetLogger().Error("Token Failed. user[" + userId + "][" + username + "][" + absPath + "]")
		return
	}

	status, msg := service.
		NewApiFileService().
		ValidDirByUserOrder(userId, absPath)

	dirAuthReturn.Status = status
	dirAuthReturn.Msg = msg
	return
}

//todo 获取文件基本信息地址
func (this *MPDCDSProServiceImpl) File(ctx context.Context, fileInfo *thriftcore.FileInfo) (r *thriftcore.FileReturn, err error) {

	token := fileInfo.Token
	absPath := fileInfo.AbsPath
	fileName := fileInfo.FileName

	r = thriftcore.NewFileReturn()
	//验证token是否有效
	m := make(map[string]string)
	isValid, err := utils.VerifyToken(m, token)
	//合法
	userIdLog := zap.String("userId", m["id"])
	userNameLog := zap.String("username", m["username"])
	accessPathLog := zap.String("accessPath", absPath)
	if isValid {
		logger.GetLogger().Info("user access", userIdLog, userNameLog, accessPathLog)
	} else {
		r.Status = -1
		r.Msg = "User authentication failed"
		logger.GetLogger().Error("User authentication failed", userIdLog, userNameLog, accessPathLog)
		return
	}

	apiFileService := service.NewApiFileService()
	data := apiFileService.GetFileInfoByAbsDir(absPath, fileName)
	if data == nil {
		r.Status = -1
		r.Msg = "File address no exist"
	} else {
		r.Status = 0
		//设置用户信息表userId

		r.Data = data
	}
	return
}

//todo 记录下载文件信息
func (this *MPDCDSProServiceImpl) SaveDownload(ctx context.Context, saveDownloadInfo *thriftcore.SaveDownloadInfo) (r *thriftcore.SaveDownloadReturn, err error) {
	token := saveDownloadInfo.Token

	//验证token是否有效
	m := make(map[string]string)
	isValid, err := utils.VerifyToken(m, token)
	//合法
	userId := m["id"]
	userIdLog := zap.String("userId", userId)
	userNameLog := zap.String("username", m["username"])
	if isValid {
		logger.GetLogger().Info("user access", userIdLog, userNameLog)
	} else {
		r.Status = -1
		r.Msg = "User authentication failed"
		logger.GetLogger().Error("User authentication failed", userIdLog, userNameLog)
		return
	}
	apiFileService := service.NewApiFileService()
	id, err := apiFileService.SaveDownLoadFileInfo(saveDownloadInfo, userId)

	if err != nil {
		logger.GetLogger().Error("SaveDownFileInfo failed", zap.String("SaveDownFileInfo", err.Error()))
		return
	}
	r = thriftcore.NewSaveDownloadReturn()
	if id != "" {
		r.Status = 0
	} else {
		r.Status = -1
		r.Msg = "SaveDownFileInfo failed"
	}
	return
}

//todo 记录采集信息
func (this *MPDCDSProServiceImpl) SaveCollect(ctx context.Context, saveCollectInfo *thriftcore.SaveCollectInfo) (r *thriftcore.SaveCollectReturn, err error) {
	panic("implement me")
}

//启动thrift server服务
func InitThriftServer() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(conf.Sysconfig.ThriftHost + ":" + conf.Sysconfig.ThriftPort)

	logger.GetLogger().Debug("thrift server start.......")
	if err != nil {
		logger.GetLogger().Error(err.Error())
	}
	handler := &MPDCDSProServiceImpl{}
	processor := thriftcore.NewMPDCDSProServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	serverTransport.Addr().String()
	logger.GetLogger().Info("thrift server in " + conf.Sysconfig.ThriftHost + ":" + conf.Sysconfig.ThriftPort)
	server.Serve()
}
