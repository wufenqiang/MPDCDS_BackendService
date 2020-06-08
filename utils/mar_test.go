package utils

import (
	"MPDCDS_BackendService/models"
	"fmt"
	"testing"
)

func TestMarshal1(t *testing.T) {
	s0 := UnMarshal(models.ApiDataInfo{})
	fmt.Println(s0)

	s1 := Marshal(s0)
	fmt.Println(s1)
}
