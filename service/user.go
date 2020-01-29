package service

import (
	"context"
	"user_center/common"
	"user_center/dao"
	gencode "user_center/protos/gencode"
)

func CreateUser(ctx context.Context, req *gencode.CreateUserRequest) (res *gencode.CreateUserResponse, err error) {
	res = &gencode.CreateUserResponse{}
	user := &dao.User{
		UID:      0,
		Name:     req.EMail,
		RealName: "",
		Role:     0,
		Phone:    req.Phone,
		EMail:    req.EMail,
		College:  req.College,
	}
	if err := dao.CreateUser(ctx,user);err !=nil{
		res.Basic = common.BadCommonResponse
		return res,nil
	}
	res.User = &gencode.UserInfo{
		Uid:                  user.UID,
		Name:                 user.Name,
		Phone:                user.Phone,
		College:              user.College,
		EMail:                user.EMail,
	}
	res.Basic = common.SuccessCommonResponse
	return res,nil
}
