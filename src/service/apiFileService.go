package service

import (
	"MPDCDS_BackendService/src/logger"
	"MPDCDS_BackendService/src/repo"
	"MPDCDS_BackendService/src/utils"
	"fmt"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/thrift/thriftcore"
	"strconv"

	proutils "gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/utils"

	"strings"
)

type ApiFileService interface {
	//根据用户ID和路径获取目录下子目录或者文件的列表信息
	GetFileByPath(userId string, dirPath string) (resMap []map[string]string)

	//根据用户ID验证路径参数是否合法，即是否有读此目录参数的权限
	ValidDirByUserOrder(userId string, absPath string) (status int16, msg string)

	//根据dirId获取当前文件地址
	GetFileInfoByAbsDir(absPath string, fileName string) (resMap map[string]string)

	//保存下载数据文件信息
	SaveDownLoadFileInfo(apiDownLoad *thriftcore.SaveDownloadInfo, userId string) (id string, error error)
}

func NewApiFileService() ApiFileService {
	return &apiFileService{}
}

type apiFileService struct {
}

func (a *apiFileService) ValidDirByUserOrder(userId, absPath string) (status int16, msg string) {

	if absPath == "" {
		status = 101
		msg = thriftcore.DirAuthReturnCodeMap[status]
		return
	} else {
		absPath = strings.ReplaceAll(absPath, "\\", "/")
		dirpaths := strings.Split(absPath, "/")

		dirpaths = proutils.ArrayFilter(dirpaths, func(dirpath string) bool {
			return dirpath != ""
		})
		lendirpaths := len(dirpaths)
		if lendirpaths == 0 {
			status = 0
			msg = thriftcore.DirAuthReturnCodeMap[status]
			return
		} else if lendirpaths == 1 {
			orderno := dirpaths[0]
			if repo.ApiOrderRepo.ValidOrderByUserId(userId, orderno) {
				status = 0
				msg = thriftcore.DirAuthReturnCodeMap[status]
				return
			} else {
				status = 2
				msg = thriftcore.DirAuthReturnCodeMap[status]
				return
			}
		} else if lendirpaths == 2 {
			orderno := dirpaths[0]
			datacode := dirpaths[1]
			if repo.ApiOrderRepo.ValidDataByUserIdOrderNo(userId, orderno, datacode) {
				status = 0
				msg = thriftcore.DirAuthReturnCodeMap[status]
				return
			} else {
				status = 3
				msg = thriftcore.DirAuthReturnCodeMap[status]
				return
			}
		} else {
			orderno := dirpaths[0]
			datacode := dirpaths[1]
			if repo.ApiOrderRepo.ValidDataByUserIdOrderNo(userId, orderno, datacode) {
				filedirpaths := dirpaths[2:]

				fmt.Println(filedirpaths)

				return
			} else {
				status = 3
				msg = thriftcore.DirAuthReturnCodeMap[status]
				return
			}
		}

		////***************************************************************************************
		////根据参数路径查询当前目录信息，判断是否为末级目录（file_index_name有值则为末级目录）
		//apiDirectory := repo.ApiDirectoryRepo.GetDirByCurrentPath(absPath)
		//if apiDirectory.Id == "" {
		//	status = 102
		//	msg = "failed: CreateFile " + absPath + ": The system cannot find the file specified."
		//	return
		//}
		//
		////根据UserId获取该用户已被授权并且有效的订单
		//
		//ApiOrders := repo.ApiOrderRepo.GetOrderByUserId(userId)
		//
		//var orderIds []interface{}
		//for _, ApiOrder := range ApiOrders {
		//	orderIds = append(orderIds, ApiOrder.OrderId)
		//}
		//
		////根据订单ID获取数据类型ID
		//datacodes := repo.ApiOrderRepo.GetDataCoreByOrderId(orderIds)
		//
		//////根据ID获取数据类型
		////apiDataInfos := repo.ApiDataInfoRepo.GetDataCodeByDataId(accessIds)
		//
		//split := strings.Split(absPath, "/")
		//s := split[1]
		//for _, e := range datacodes {
		//	if e == s {
		//		status = 0
		//		msg = "Path validation successful"
		//		return
		//	}
		//}
		//
		//status = 103
		//msg = "Path validation Failed"
		//return
	}
}

func (a *apiFileService) GetFileByPath(userId, dirPath string) (resMap []map[string]string) {
	dirPath = strings.ReplaceAll(dirPath, "\\", "/")
	dirpaths := strings.Split(dirPath, "/")

	dirpaths = proutils.ArrayFilter(dirpaths, func(dirpath string) bool {
		return dirpath != ""
	})

	if dirpaths == nil {
		//如果目前为"/"则直接返回用户订阅的数据类型的编码作为一级目录

		//根据UserId获取该用户已被授权并且有效的订单
		apiOrders := repo.ApiOrderRepo.GetOrderByUserId(userId)
		for _, apiOrder := range apiOrders {
			var tmpMap map[string]string
			tmpMap = make(map[string]string)
			tmpMap["Size"] = "0"
			tmpMap["Type"] = "dir"
			tmpMap["FileName"] = dirPath + apiOrder.OrderNo
			tmpMap["Modify"] = ""
			resMap = append(resMap, tmpMap)

		}
		return
	} else {
		lendirpaths := len(dirpaths)
		if lendirpaths == 1 {
			orderno := dirpaths[0]
			datalist := repo.ApiOrderRepo.GetDataCoreByOrderNo(orderno)
			for _, datacode := range datalist {
				var tmpMap map[string]string
				tmpMap = make(map[string]string)
				tmpMap["Size"] = "0"
				tmpMap["Type"] = "dir"
				tmpMap["FileName"] = dirPath + "/" + datacode
				tmpMap["Modify"] = ""
				resMap = append(resMap, tmpMap)
			}
			//}else if(lendirpaths==2){
			//	orderno:=dirpaths[0]
			//	datacode:=dirpaths[1]
			//
			//}else{
			//	orderno:=dirpaths[0]
			//	datacode:=dirpaths[1]
		}

		//根据UserId获取该用户已被授权并且有效的订单
		apiOrders := repo.ApiOrderRepo.GetOrderByUserId(userId)
		for _, apiOrder := range apiOrders {
			orderno := apiOrder.OrderNo
			if dirPath == "/"+orderno {
				orderid := apiOrder.OrderId

				var orderIds []interface{}
				orderIds = append(orderIds, orderid)
				//根据订单ID获取数据类型ID
				accessIds := repo.ApiOrderRepo.GetDataCoreByOrderId(orderIds)

				for _, accessId := range accessIds {
					var tmpMap map[string]string
					tmpMap = make(map[string]string)
					tmpMap["Size"] = "0"
					tmpMap["Type"] = "dir"
					tmpMap["FileName"] = dirPath + "/" + accessId
					tmpMap["Modify"] = ""
					resMap = append(resMap, tmpMap)
				}

				return
			}
		}

		if true {
			//根据ID获取数据类型
			//apiDataInfos := repo.ApiDataInfoRepository.GetDataCodeByDataId(accessIds)

			//根据参数路径查询当前目录信息，判断是否为末级目录（file_index_name有值则为末级目录）
			currentDirectory := repo.ApiDirectoryRepo.GetDirByCurrentPath(dirPath)
			if currentDirectory.FileIndexName == "" {
				//根据路径参数查询需要返回的目录或者文件信息
				apiDirectorys := repo.ApiDirectoryRepo.GetDirByParentPath(dirPath)
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

				pathCode, err := utils.AesEcryptStr(currentDirectory.CurrentDir)
				if err != nil {
					logger.GetLogger().Error("字符窜加密失败！")
				}

				apiFiles := repo.ApiFileRepo.GetFileByIndexNameAndDirId(currentDirectory.FileIndexName, pathCode)
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
		} else {

		}
		return
	}

}

func (a *apiFileService) GetFileInfoByAbsDir(absPath, fileName string) (resMap map[string]string) {

	//根据当前目录绝对地址从目录表中获取file_index_name
	apiDirectory := repo.ApiDirectoryRepo.GetDirByCurrentPath(absPath)
	file_index_name := apiDirectory.FileIndexName

	if file_index_name == "" {
		file_index_name = proutils.UnMarshal(models.ApiFile{})
	}

	pathCode, err := utils.AesEcryptStr(absPath)
	if err != nil {
		logger.GetLogger().Error("字符窜加密失败！")
	}

	//根据file_index_name从文件信息表查询文件真实地址
	apiFile := repo.ApiFileRepo.GetFileByIndexNameAndDirIdAndFileName(file_index_name, pathCode, fileName)
	if apiFile.Id != "" {
		resMap = make(map[string]string)
		resMap["file_address"] = apiFile.FileAddress
		resMap["access_id"] = apiFile.AccessId
		resMap["file_id"] = apiFile.Id
		return
	}
	return
}

func (a *apiFileService) SaveDownLoadFileInfo(apiDownLoad *thriftcore.SaveDownloadInfo, userId string) (id string, error error) {
	id, err := repo.ApiDownRepo.SaveDownLoadFileInfo(apiDownLoad, userId)
	return id, err
}
