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
	res,err := service.CreateUser(ctx,req)
	return nil,nil
}
func (server *UserServer)GetUserByID(ctx context.Context,req *gencode.GetUserInfoByIDRequest)(res *gencode.GetUserInfoByIDResponse,err error){
	return nil,nil
}
func (server *UserServer)SayHello(ctx context.Context,req *gencode.HelloRequest)(res *gencode.HelloResponse,err error){
	logrus.Debugf("hello form : %v",req.Name)
	res = &gencode.HelloResponse{Name: "benji's server"}
	return res,nil
}
func (server *UserServer)VerifyUserIdentity(context.Context, *gencode.VerifyUserIdentityRequest) (res *gencode.VerifyUserIdentityResponse,err error){

}
