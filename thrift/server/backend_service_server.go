package server

import (
	"MPDCDS_BackendService/conf"
	"MPDCDS_BackendService/thrift/MPDCDS_BackendService"
	"MPDCDS_BackendService/utils"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type MPDCDS_BackendServiceImpl struct {
}

func (this *MPDCDS_BackendServiceImpl) Auth(ctx context.Context, user string, password string) (r string, err error) {
	//todo 查询es验证用户名和密码是否合法

	if true {
		//用户名和userId合法,使用jwt生成token,userId为es中查询的用户信息主键
		r, err = utils.GenerateToken(user, "userId")
		if err != nil {
			fmt.Print("生成token失败！")
		}
	}
	return
}

func (this *MPDCDS_BackendServiceImpl) Lists(ctx context.Context, token string, pwd string) (r []map[string]string, err error) {
	//验证token是否有效
	m := make(map[string]string)
	isValid, err := utils.VerifyToken(m, token)
	//合法
	if isValid {
		//todo 根据用户信息、当前目录从es中查询数据目录和数据列表
		fmt.Print("用户id:" + m["id"] + "=========用户名称" + m["username"])
	}
	return
}

func (this *MPDCDS_BackendServiceImpl) File(ctx context.Context, pwd string, path string) (r map[string]string, err error) {
	//todo 根据当前目录和文件名称查询文件真实地址

	return
}

//启动thrift server服务
func InitMpdcdsBackendServiceServer() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(conf.Sysconfig.NetworkAddr)
	fmt.Println("thrift server start.......")
	if err != nil {
		fmt.Println("Error!", err)
	}
	handler := &MPDCDS_BackendServiceImpl{}
	processor := MPDCDS_BackendService.NewMPDCDS_BackendServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", conf.Sysconfig.NetworkAddr)
	server.Serve()
}
