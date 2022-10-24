package main;

import (
	"log"
	"net/rpc"
	s "rpcpoc/service"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":50001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &s.Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}