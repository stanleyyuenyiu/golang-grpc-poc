package main

import (
   "time"
   "flag"
   "fmt"
   udp "udppoc/lib/udp"
	  c "udppoc/lib/struct/data"
)


var (
	port = flag.Int("port", 6001, "The server port")
  addr = flag.String("addr", "255.255.255.255", "The server addr")
)

func main() {
  
  sendCh := make(chan c.CommData)
	connStatus, msg := udp.Send(sendCh, *addr, *port, 6000)
	go RunPrintConn(connStatus)
	go RunPrintMsg(msg)

  time.Sleep(1 * time.Second)
  count := 0
	for {
    msg := udp.BuildMsg("127.0.0.1", fmt.Sprintf("MSGID-%v", count), "counter", count)
    sendCh <- *msg
    time.Sleep(2 * time.Second)
    count += 1
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
