package controllers

import (
	"github.com/kataras/iris"
	"goserver-api/models/bak"
	"goserver-api/service"
	"log"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: service.NewUserServices()}
}

func (g *UserController) PostLogin() bak.Result {
	var m map[string]string
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	result := g.Service.Login(m)
	return result
}

func (g *UserController) PostSave() (result bak.Result) {
	var user bak.User
	if err := g.Ctx.ReadJSON(&user); err != nil {
		log.Println(err)
		result.Msg = "数据错误"
		return
	}

	return g.Service.Save(user)
}
