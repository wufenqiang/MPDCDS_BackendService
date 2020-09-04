package test

import (
	"fmt"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"
	proutils "gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/utils"
	"testing"
)

func TestStructToJsonDemo(t *testing.T) {

	apiuser := models.Api_user{
		Id:         0,
		Username:   "test",
		Password:   "",
		CreateTime: proutils.Now2TimeString(),
		Realname:   "",
		Phone:      "",
		Email:      "",
		Appkey:     "",
		Status:     "",
		Remark:     "",
		DelFlag:    "",
		DeptId:     0,
	}

	jsonstr, _ := proutils.Struct2JsonStr(apiuser)
	fmt.Println(jsonstr)
}
