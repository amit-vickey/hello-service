package server

import (
	"context"
	pb "github.com/amit/hello-service/rpc"
)

type HelloWorldServer struct{}

func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Text: "Hello " + req.Subject}, nil
}
