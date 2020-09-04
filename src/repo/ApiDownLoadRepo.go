package repo

import (
	"MPDCDS_BackendService/src/conf"
	"MPDCDS_BackendService/src/datasource"
	"MPDCDS_BackendService/src/logger"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/thrift/thriftcore"

	"MPDCDS_BackendService/src/utils"
	"context"
	proutils "gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/utils"
	"go.uber.org/zap"
	"time"
)

var ApiDownRepo = new(apiDownRepo)

type apiDownRepo struct{}

//保存下载数据文件信息
func (a *apiDownRepo) SaveDownLoadFileInfo(apiDownLoad *thriftcore.SaveDownloadInfo, userId string) (id string, error error) {
	dataClient := datasource.ES.GetDB()

	/*需要转换的时间类型字符串*/
	startTime, _ := time.ParseInLocation(conf.Layout, apiDownLoad.StartTime, conf.Loc)
	endTime, _ := time.ParseInLocation(conf.Layout, apiDownLoad.EndTime, conf.Loc)
	apiDownLoad_model := models.ApiDownload{"", apiDownLoad.AccessID, apiDownLoad.FileID, startTime, endTime, userId, time.Now()}

	res, err := dataClient.Index().
		Index(proutils.UnMarshal(models.ApiDownload{})).
		Id(utils.Uuid()).
		BodyJson(apiDownLoad_model).
		Do(context.Background())
	if err != nil {
		logger.GetLogger().Error("保存下载文件信息失败！", zap.String("error", err.Error()))
		return "", err
	}
	return res.Id, err
}
