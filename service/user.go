package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"user_center/common"
	"user_center/dao"
	gencode "user_center/protos/gencode"
)

func CreateUser(ctx context.Context, req *gencode.CreateUserRequest) (res *gencode.CreateUserResponse, err error) {
	res = &gencode.CreateUserResponse{}
	user := &dao.User{
		Name:     req.EMail,
		Role:     0,
		Phone:    req.Phone,
		EMail:    req.EMail,
		College:  req.College,
	}
	if err := dao.CreateUser(ctx,user);err !=nil{
		res.Basic = common.BadCommonResponse
		return res,nil
	}
	res.User = ModelUserToUserInfo(user)
	res.Basic = common.SuccessCommonResponse
	return res,nil
}
func VerifyUser(ctx context.Context,req *gencode.VerifyUserIdentityRequest)(res *gencode.VerifyUserIdentityResponse,err error){
	res = &gencode.VerifyUserIdentityResponse{}
	user,err  := dao.GetUserInfoByName(ctx,req.UserName)
	if err != nil{
		logrus.Errorf("fail to get userInfoByName,err:%v",err)
		res.Basic = common.NotFoundCommonResponse
		return res,err
	}
	res.Basic = common.SuccessCommonResponse
	res.User = ModelUserToUserInfo(user)
	return res,nil

}
func GetUserInfoById(ctx context.Context,req *gencode.GetUserInfoByIDRequest)(res *gencode.GetUserInfoByIDResponse,err error){
	res = &gencode.GetUserInfoByIDResponse{}
	user,err := dao.GetUserInfoById(ctx,req.Uid)
	if err != nil{
		logrus.Errorf("fail to get userInfoById,err:%v",err)
		res.Basic = common.NotFoundCommonResponse
		return res,err
	}
	res.Basic = common.SuccessCommonResponse
	res.User = ModelUserToUserInfo(user)
	return res,nil
}
func ModelUserToUserInfo(user *dao.User)*gencode.UserInfo{
	userInfo := &gencode.UserInfo{
		Uid:                  user.ID,
		Name:                 user.Name,
		Phone:                user.Phone,
		College:              user.College,
		EMail:                user.EMail,
	}
	return userInfo
}
