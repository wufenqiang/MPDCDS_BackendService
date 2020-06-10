package service

import (
	"MPDCDS_BackendService/models"
	"MPDCDS_BackendService/repo"
	"MPDCDS_BackendService/utils"
	"strconv"
	"strings"
)

type ApiFileService interface {
	//根据用户ID和路径获取目录下子目录或者文件的列表信息
	GetFileByPath(userId, dirPath string) (resMap []map[string]string)

	//根据用户ID验证路径参数是否合法，即是否有读此目录参数的权限
	ValidDirByUserOrder(userId, absPath string) (status int16, msg string)

	//根据dirId获取当前文件地址
	GetFileInfoByAbsDir(absPath, fileName string) (r map[string]string)
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

func (a apiFileService) ValidDirByUserOrder(userId, absPath string) (status int16, msg string) {

	if absPath == "" {
		status = 101
		msg = "Parameter cannot be empty"
		return
	}

	if absPath == "/" {
		status = 0
		msg = "Path validation successful"
		return
	}

	//根据参数路径查询当前目录信息，判断是否为末级目录（file_index_name有值则为末级目录）
	currentDirectory := apiDirectoryRepository.GetDirByCurrentPath(absPath)
	if currentDirectory.Id == "" {
		status = 102
		msg = "failed: CreateFile " + absPath + ": The system cannot find the file specified."
		return
	}

	//根据UserId获取该用户已被授权并且有效的订单
	apiOrders := apiOrderRepository.GetOrderByUserId(userId)

	var orderIds []interface{}
	for _, e := range apiOrders {
		orderIds = append(orderIds, e.Id)
	}

	//根据订单ID获取数据类型ID
	accessIds := apiDataOrderShipRepository.GetDataOrderShipListByOrderId(orderIds)

	//根据ID获取数据类型
	apiDataInfos := apiDataInfoRepository.GetApiDataInfoById(accessIds)

	split := strings.Split(absPath, "/")
	s := split[1]
	for _, e := range apiDataInfos {
		if e.DataCode == s {
			status = 0
			msg = "Path validation successful"
			return
		}
	}

	status = 103
	msg = "Path validation Failed"
	return
}

func (a apiFileService) GetFileByPath(userId, dirPath string) (resMap []map[string]string) {

	//根据UserId获取该用户已被授权并且有效的订单
	apiOrders := apiOrderRepository.GetOrderByUserId(userId)

	var orderIds []interface{}
	for _, e := range apiOrders {
		orderIds = append(orderIds, e.Id)
	}

	//根据订单ID获取数据类型ID
	accessIds := apiDataOrderShipRepository.GetDataOrderShipListByOrderId(orderIds)

	//根据ID获取数据类型
	apiDataInfos := apiDataInfoRepository.GetApiDataInfoById(accessIds)

	//如果目前为"/"则直接返回用户订阅的数据类型的编码作为一级目录
	if dirPath == "/" {
		for _, e := range apiDataInfos {
			var tmpMap map[string]string
			tmpMap = make(map[string]string)
			tmpMap["Size"] = "0"
			tmpMap["Type"] = "dir"
			tmpMap["FileName"] = "/" + e.DataCode
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
			tmpMap = make(map[string]string)
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
			tmpMap = make(map[string]string)
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

func (a apiFileService) GetFileInfoByAbsDir(absPath, fileName string) (r map[string]string) {

	//根据当前目录绝对地址获取从目录表中获取file_index_name
	apiDirectory := apiDirectoryRepository.GetDirByCurrentPath(absPath)
	file_index_name := apiDirectory.FileIndexName

	if file_index_name == "" {
		file_index_name = utils.UnMarshal(models.ApiFile{})
	}
	//根据file_index_name从文件信息表查询文件真实地址
	apiFile := apiFileRepository.GetFileByIndexNameAndDirIdAndFileName(file_index_name, utils.AesEcryptStr(absPath), fileName)

	if apiFile.Id != "" {
		r = make(map[string]string)
		r["file_address"] = apiFile.FileAddress
		return
	}
	return
}
