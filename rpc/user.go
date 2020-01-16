package rpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"user_center/protos/gencode"
)
type UserServer struct {
	gencode.UnimplementedUserServer
}
func (server *UserServer)CreateUser(ctx context.Context,req *gencode.CreateUserRequest)(res *gencode.CreateUserResponse,err error){
	return nil,nil
}
func (server *UserServer)GetUserByID(ctx context.Context,req *gencode.GetUserByIDRequest)(res *gencode.GetUserByIDResponse,err error){
	return nil,nil
}
func (server *UserServer)SayHello(ctx context.Context,req *gencode.HelloRequest)(res *gencode.HelloResponse,err error){
	logrus.Debugf("hello form : %v",req.Name)
	res = &gencode.HelloResponse{Name: "benji's server"}
	return res,nil
}