package service;

import (
	"errors"
	"log"
)


type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	log.Printf("[Divide]  Receive: %v %v", args.A, args.B)
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	log.Printf("[Divide]  Receive: %v %v", args.A, args.B)
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}