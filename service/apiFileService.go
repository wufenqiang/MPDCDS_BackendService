package service

import (
	"MPDCDS_BackendService/repo"
	"MPDCDS_BackendService/utils"
	"fmt"
	"strconv"
)

type ApiFileService interface {
	GetFileByPath(userId, dirPath string) (resMap []map[string]string)
}

func NewApiFileService() ApiFileService {
	return &apiFileService{}
}

type apiFileService struct {
}

var (
	apiDirectoryRepository     = repo.NewApiDirectoryRepository()
	apiFileRepository          = repo.NewApiFileRepository()
	apiOrderRepository         = repo.NewApiOrderRepository()
	apiDataOrderShipRepository = repo.NewApiDataOrderShipRepository()
	apiDataInfoRepository      = repo.NewApiDataInfoRepository()
)

func (a apiFileService) GetFileByPath(userId, dirPath string) (resMap []map[string]string) {

	//根据UserId获取该用户已被授权并且有效的订单
	apiOrders := apiOrderRepository.GetOrderByUserId(userId)
	fmt.Println("apiOrders", apiOrders)

	var orderIds []interface{}
	for _, e := range apiOrders {
		orderIds = append(orderIds, e.Id)
	}

	//根据订单ID获取数据类型ID
	accessIds := apiDataOrderShipRepository.GetDataOrderShipListByOrderId(orderIds)
	fmt.Println("accessIds", accessIds)

	//根据ID获取数据类型
	apiDataInfos := apiDataInfoRepository.GetApiDataInfoById(accessIds)
	fmt.Println("apiDataInfos", apiDataInfos)

	//如果目前为"/"则直接返回用户订阅的数据类型的编码作为一级目录
	if dirPath == "/" {
		for _, e := range apiDataInfos {
			var tmpMap map[string]string
			tmpMap["Size"] = "0"
			tmpMap["Type"] = "dir"
			tmpMap["FileName"] = e.DataCode
			tmpMap["Modify"] = ""
			resMap = append(resMap, tmpMap)
		}
		return
	}

	//根据参数路径查询当前目录信息，判断是否为末级目录（file_index_name有值则为末级目录）
	currentDirectory := apiDirectoryRepository.GetDirByCurrentPath(dirPath)
	if currentDirectory.FileIndexName == "" {
		//根据路径参数查询需要返回的目录或者文件信息
		apiDirectorys := apiDirectoryRepository.GetDirByParentPath(dirPath)
		for _, e := range apiDirectorys {
			var tmpMap map[string]string
			tmpMap["Size"] = "0"
			tmpMap["Type"] = "dir"
			tmpMap["FileName"] = e.CurrentDir
			tmpMap["Modify"] = ""
			resMap = append(resMap, tmpMap)
		}
		return
	} else {
		//根据api_directory目录中的file_index_name索引名和唯一标识查询当前目录下的文件信息
		apiFiles := apiFileRepository.GetFileByIndexNameAndDirId(currentDirectory.FileIndexName, utils.AesEcryptStr(currentDirectory.CurrentDir))
		for _, e := range apiFiles {
			var tmpMap map[string]string
			tmpMap["Size"] = strconv.FormatInt(e.FileSize, 10)
			tmpMap["Type"] = "file"
			tmpMap["FileName"] = e.FileName
			tmpMap["Modify"] = e.CreateTime.Format("2006-01-02 15:04:05")
			resMap = append(resMap, tmpMap)
		}
		return
	}

	return
}
