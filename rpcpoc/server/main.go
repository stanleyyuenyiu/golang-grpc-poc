package main;

import (
	"net"
	"log"
	"net/rpc"
	"net/http"
	s "rpcpoc/service"
)

func main() {
	arith := new(s.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	lis, e := net.Listen("tcp", ":50001")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Printf("server listening at %v", lis.Addr())
	http.Serve(lis, nil)
}