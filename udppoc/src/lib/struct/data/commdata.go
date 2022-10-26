package data

import (
	"fmt"
	"net"
)


type CommData struct {
	Identifier string
	AddrPort net.UDPAddr
	SenderIP	string
	ReceiverIP	string
	MsgID string
	DataType string
	DataValue any
}

func (data *CommData) PrintData() {
	fmt.Println("=== Communication data ===")
	fmt.Println("Identifier:", data.Identifier)
	fmt.Println("SenderIP:", data.SenderIP)
	fmt.Println("ReceiverIP:", data.ReceiverIP)
	fmt.Println("Message ID:", data.MsgID)
	fmt.Println("DataType:", data.DataType)
	fmt.Println("DataValue:", data.DataValue)
}