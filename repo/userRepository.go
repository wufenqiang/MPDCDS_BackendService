package repo

import (
	esdatasource "MPDCDS_BackendService/datasource/elasticsearch"
	"MPDCDS_BackendService/datasource/mysql"
	"MPDCDS_BackendService/logger"
	"MPDCDS_BackendService/models"
	"MPDCDS_BackendService/models/bak"
	"MPDCDS_BackendService/utils"
	"context"
	"github.com/olivere/elastic/v7"
	"reflect"

	// "github.com/spf13/cast"
	"log"
)

type UserRepository interface {
	GetUserByUserNameAndPwd(username string, password string) (user bak.User)
	GetUserByUsername(username string) (user bak.User)
	Save(user bak.User) (int, bak.User)
	GetUserByName(username string) (user models.Hf_platform_user)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

//登录
func (n userRepository) GetUserByUserNameAndPwd(username string, password string) (user bak.User) {
	db := mysql.GetDB()
	db.Where("username = ? and password = ?", username, password).First(&user)

	return
}
func (n userRepository) GetUserByUsername(username string) (user bak.User) {
	db := mysql.GetDB()
	db.Where("username = ?", username).First(&user)
	return
}

//添加/修改
func (n userRepository) Save(user bak.User) (int, bak.User) {
	code := 0
	tx := mysql.GetDB().Begin()
	defer utils.Defer(tx, &code)
	if user.ID != 0 {
		var oldUser bak.User
		mysql.GetDB().First(&oldUser, user.ID)
		user.CreatedAt = oldUser.CreatedAt
		user.Username = oldUser.Username
		if user.Name == "" {
			user.Name = oldUser.Name
		}
		if user.Email == "" {
			user.Email = oldUser.Email
		}
		if user.Mobile == "" {
			user.Mobile = oldUser.Mobile
		}
		if user.QQ == "" {
			user.QQ = oldUser.QQ
		}
		if user.Gender == 0 {
			user.Gender = oldUser.Gender
		}
		if user.Age == 0 {
			user.Age = oldUser.Age
		}
		if user.Remark == "" {
			user.Remark = oldUser.Remark
		}
	}
	if user.Password != "" {
		user.Password = utils.GetMD5String(user.Password)
	}
	if err := tx.Save(&user).Error; err != nil {
		log.Println(err)
		code = -1
	}
	return code, user
}

func (n userRepository) PostLogin() bool {
	return false
}

func (n userRepository) GetUserByName(userName string) models.Hf_platform_user {
	esClient := esdatasource.GetESClient()

	q := elastic.NewQueryStringQuery("Username:" + userName)
	res, err := esClient.Search("hf_platform_user").
		Size(1).
		From(0).
		Query(q).Do(context.Background())

	if err != nil {
		logger.GetLogger().Error("GetUserByName")
	}

	var user models.Hf_platform_user
	for _, item := range res.Each(reflect.TypeOf(user)) { //从搜索结果中取数据的方法
		user = item.(models.Hf_platform_user)
	}

	return user
}
