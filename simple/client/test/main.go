package client_test;

import (
	"log"
)
type Interface interface {
	Test()    // remove and return element Len() - 1.
}

func Init(h Interface) {
	log.Printf("%p", h)
}

func Test(h Interface)  {
	 h.Test()
}