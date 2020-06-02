package service

import (
	"goserver-api/middleware"
	"goserver-api/models/bak"
	"goserver-api/repo/mysql"
	"goserver-api/utils"
	// "fmt"
	// "github.com/spf13/cast"
	// "log"
)

type UserService interface {
	Login(m map[string]string) (result bak.Result)
	Save(user bak.User) (result bak.Result)
}
type userServices struct {
}

func NewUserServices() UserService {
	return &userServices{}
}

var userRepo = mysql.NewUserRepository()

/*
登录
*/
func (u userServices) Login(m map[string]string) (result bak.Result) {

	if m["username"] == "" {
		result.Code = -1
		result.Msg = "请输入用户名！"
		return
	}
	if m["password"] == "" {
		result.Code = -1
		result.Msg = "请输入密码！"
		return
	}
	user := userRepo.GetUserByUserNameAndPwd(m["username"], utils.GetMD5String(m["password"]))
	if user.ID == 0 {
		result.Code = -1
		result.Msg = "用户名或密码错误!"
		return
	}
	user.Token = middleware.GenerateToken(user)
	result.Code = 0
	result.Data = user
	result.Msg = "登录成功"
	return
}

/*
保存
*/
func (u userServices) Save(user bak.User) (result bak.Result) {
	//添加
	if user.ID == 0 {
		agen := userRepo.GetUserByUsername(user.Username)
		if agen.ID != 0 {
			result.Msg = "登录名重复,保存失败"
			return
		}
	}
	code, p := userRepo.Save(user)
	if code == -1 {
		result.Code = -1
		result.Msg = "保存失败"
		return
	}
	result.Code = 0
	result.Data = p
	return
}
