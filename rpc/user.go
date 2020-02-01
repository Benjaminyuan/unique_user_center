package rpc

import (
	"context"
	"github.com/sirupsen/logrus"
	gencode "user_center/protos/gencode"
	"user_center/service"
)
type UserServer struct {
	gencode.UnimplementedUserServer
}
func (server *UserServer)CreateUser(ctx context.Context,req *gencode.CreateUserRequest)(res *gencode.CreateUserResponse,err error){
	logrus.Debugf("CreateUser req:%v",req)
	res,err  = service.CreateUser(ctx,req)
	return res,err
}
func (server *UserServer)GetUserInfoByID(ctx context.Context,req *gencode.GetUserInfoByIDRequest)(res *gencode.GetUserInfoByIDResponse,err error){
	logrus.Debugf("GetUserInfoByID req:%v",req)
	return nil,nil
}
func (server *UserServer)SayHello(ctx context.Context,req *gencode.HelloRequest)(res *gencode.HelloResponse,err error){
	logrus.Debugf("hello form : %v",req.Name)
	res = &gencode.HelloResponse{Name: "benji's server"}
	return res,nil
}
func (server *UserServer)VerifyUserIdentity(ctx context.Context,req  *gencode.VerifyUserIdentityRequest) (res *gencode.VerifyUserIdentityResponse,err error){
	logrus.Debugf("VerifyUserIdentity req:%v",req)
	res,err = service.VerifyUser(ctx,req)
	return res,err
}
