package simple_service;

import (
	"context"
	"log"
	"google.golang.org/grpc"
	pb "simple/protos/simple"
) ;

type SimpleService struct {
	pb.UnimplementedCallServer
}

func (s *SimpleService) ReqRes(ctx context.Context, in *pb.CallRequest) (*pb.CallResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.CallResponse{Message: "Hello " + in.GetName()}, nil
}

func Register(s grpc.ServiceRegistrar) {
	pb.RegisterCallServer(s, &SimpleService{});
}