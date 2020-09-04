package test

import (
	"MPDCDS_BackendService/src/datasource"
	"fmt"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"
	proutils "gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/utils"
	"testing"
)

//测试mysql转换对象
func TestFind(t *testing.T) {
	user := "123456"
	//password :=""

	mysqlclient := datasource.MySQL.GetDB()

	var apiuser0 models.Api_user

	fmt.Println(mysqlclient.HasTable(apiuser0))

	e := mysqlclient.Find(&apiuser0, "username = ?", user).Error

	if e == nil {

		jsonStr0, _ := proutils.Struct2JsonStr(apiuser0)
		apiuser1 := new(models.Api_user)
		proutils.JsonStr2Struct(jsonStr0, apiuser1)
		jsonStr1, _ := proutils.Struct2JsonStr(apiuser1)

		fmt.Println(apiuser0)
		fmt.Println(apiuser1)
		fmt.Println(jsonStr0)
		fmt.Println(jsonStr1)

		fmt.Println("ture")
	} else {
		fmt.Println("false")
	}

}
