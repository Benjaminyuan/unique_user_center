package rpc

import(
	"context"
	pb "user_center/server/protos/sso"
)
type UserServer struct {
	pb.UnimplementedUserServer
}
func (server *UserServer)CreateUser(ctx context.Context,req *pb.CreateUserRequest)(rep *pb.CreateUserResponse,err error){
	return nil,nil
}
func (server *UserServer)GetUserByID(ctx context.Context,req *pb.GetUserByIDRequest)(rep *pb.GetUserByIDResponse,err error){
	return nil,nil
}
