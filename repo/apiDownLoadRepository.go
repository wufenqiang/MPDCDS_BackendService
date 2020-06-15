package repo

import (
	esdatasource "MPDCDS_BackendService/datasource/elasticsearch"
	"MPDCDS_BackendService/logger"
	"MPDCDS_BackendService/models"
	"MPDCDS_BackendService/thrift/MPDCDS_BackendService"
	"MPDCDS_BackendService/utils"
	"context"
	"go.uber.org/zap"
	"time"
)

type ApiDownRepository interface {
	//保存下载数据文件信息
	SaveDownLoadFileInfo(apiDownLoad *MPDCDS_BackendService.ApiDownLoad, userId string) (id string, error error)
}

type apiDownRepository struct{}

func NewApiDownRepository() ApiDownRepository {
	return &apiDownRepository{}
}

func (a apiDownRepository) SaveDownLoadFileInfo(apiDownLoad *MPDCDS_BackendService.ApiDownLoad, userId string) (id string, error error) {
	const Layout = "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")
	/*需要转换的时间类型字符串*/
	startTime, _ := time.ParseInLocation(Layout, apiDownLoad.StartTime, loc)
	endTime, _ := time.ParseInLocation(Layout, apiDownLoad.EndTime, loc)
	apiDownLoad_model := models.ApiDownload{"", apiDownLoad.AccessID, apiDownLoad.FileID, startTime, endTime, userId, time.Now()}
	esClient := esdatasource.GetESClient()
	res, err := esClient.Index().
		Index(utils.UnMarshal(models.ApiDownload{})).
		Id(utils.Uuid()).
		BodyJson(apiDownLoad_model).
		Do(context.Background())
	if err != nil {
		logger.GetLogger().Error("保存下载文件信息失败！", zap.String("error", err.Error()))
		return "", err
	}
	return res.Id, err
}
