package listener

import (
	"flag"
	udp "udppoc/lib/udp"
	c "udppoc/lib/struct/data"
	"time"
)

var (
	port = flag.Int("port", 6001, "The server port")
)
func Run() {
	sendCh := make(chan c.CommData)
	connStatus, msg := udp.Listen(sendCh, "", *port)
	go RunPrintConn(connStatus)
	go RunPrintMsg(msg)
	for{
		time.Sleep(10 * time.Second)
	}
}

func RunPrintConn(connStatus <-chan c.ConnData) {
	for{
		data := <- connStatus
    	data.PrintData()
	}
}

func RunPrintMsg(msg <-chan c.CommData) {
	for{
		data := <- msg
    	data.PrintData()
	}
}
