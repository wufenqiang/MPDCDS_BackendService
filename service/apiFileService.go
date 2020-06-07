package service

import (
	"MPDCDS_BackendService/models"
	"MPDCDS_BackendService/repo"
	"fmt"
)

type ApiFileService interface {
	GetFileByPath(dirPath string) []models.ApiFile
}

func NewApiFileService() ApiFileService {
	return &apiFileService{}
}

type apiFileService struct {
}

var (
	apiDirectoryRepository = repo.NewApiDirectoryRepository()
	apiFileRepository      = repo.NewApiFileRepository()
)

func (a apiFileService) GetFileByPath(dirPath string) []models.ApiFile {

	apiDirectory := apiDirectoryRepository.GetDirByPath(dirPath)
	apiFiles := apiFileRepository.GetFileByDirId(apiDirectory.DataId, "")
	fmt.Println(apiFiles)
	return apiFiles
}