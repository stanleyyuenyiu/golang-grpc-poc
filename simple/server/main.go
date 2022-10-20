package main;

import (
	"fmt"
	"net"
	"log"
	"flag"
	"google.golang.org/grpc"
	simple "simple/service/simple"
	topk "simple/service/topk"
) ;

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	
	simple.Register(s)
	topk.Register(s)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



