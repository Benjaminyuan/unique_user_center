package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"user_center/conf"
	"user_center/dao"
	"user_center/rpc"
	pb "user_center/server/protos/sso"
)

func main(){
	logrus.SetLevel(logrus.DebugLevel)
	err := conf.InitConf()
	if err != nil{
		logrus.Fatalf("fail to init conf, err: %v",err)
	}
	err = dao.InitDB()
	if err != nil{
		logrus.Fatalf("fail to init DB, err: %v",err)
	}
	lis,err := net.Listen("",conf.Data.RPCConf.Addr)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s,&rpc.UserServer{})
	if err:= s.Serve(lis);err != nil{
		logrus.Fatalf("fail to  start listen, err: %v",err)
	}
}