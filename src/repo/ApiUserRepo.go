package repo

import (
	"MPDCDS_BackendService/src/datasource"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"
)

var ApiUserRepo = new(apiUserRepo)

type apiUserRepo struct{}

func (a *apiUserRepo) GetUserByUserName(user string) (models.Api_user, error) {
	dataClient := datasource.MySQL.GetDB()

	var apiuser models.Api_user

	e := dataClient.
		Find(&apiuser, "username = ?", user).
		Error

	return apiuser, e
}
