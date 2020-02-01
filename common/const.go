package common

import (
	gencode "user_center/protos/gencode"
)
var (
	SuccessCommonResponse = &gencode.CommonResponse{Code: 0, Info:"success"}
	BadCommonResponse =  &gencode.CommonResponse{Code: 1, Info:"server error"}
	NotFoundCommonResponse = &gencode.CommonResponse{Code: 2, Info:"not found"}

)