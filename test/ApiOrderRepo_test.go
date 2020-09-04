package test

import (
	"MPDCDS_BackendService/src/repo"
	"testing"
)

func TestValidDataByUserIdOrderNo(t *testing.T) {
	userId := "2"
	orderno := "2020090117225013046317664"
	datacode := "93_001_DM_CHN_VISJDF_5M_STCO"
	flag := repo.ApiOrderRepo.ValidDataByUserIdOrderNo(userId, orderno, datacode)
	println(flag)
}
func TestValidOrderByUserId(t *testing.T) {
	userId := "2"
	orderno := "2020090117225013046317664"
	flag := repo.ApiOrderRepo.ValidOrderByUserId(userId, orderno)
	println(flag)
}
