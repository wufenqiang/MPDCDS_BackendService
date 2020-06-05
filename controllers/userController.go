package controllers

import (
	"MPDCDS_BackendService/models/bak"
	"MPDCDS_BackendService/service"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
	"log"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: service.NewUserServices()}
}

func (g *UserController) GetLogin() bak.Result {
	var m map[string]string
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	result := g.Service.Login(m)
	return result
}

//func (g *UserController) PostSave() (result bak.Result) {
//	var user bak.User
//	if err := g.Ctx.ReadJSON(&user); err != nil {
//		log.Println(err)
//		result.Msg = "数据错误"
//		return
//	}
//
//	return g.Service.Save(user)
//}

func (g *UserController) GetUserByUserName() (result bak.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}

	if m["username"] == "" || m["username"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 username"
		return
	}

	result = g.Service.GetUserByName(cast.ToString(m["username"]))
	return
}
