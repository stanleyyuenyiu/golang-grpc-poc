package udp

import (
	"net"
	"log"
   "time"
   //"bytes"
   "encoding/json"
   "fmt"
   c "udppoc/lib/struct/data"
)
const com_id = "2323"
func Listen(sendCh chan c.CommData, listenAddr string,  listenPort int) (<-chan c.ConnData, <-chan c.CommData) {
  commSend := make(chan c.CommData)
  commReceive := make(chan c.CommData, 1)
  commSentStatus := make(chan c.ConnData)
  receivedMsg := make(chan c.CommData)

  go listen(commReceive, fmt.Sprintf(":%v", listenPort), false)
  go broadcast(commSend, "", fmt.Sprintf(":%v", listenPort), fmt.Sprintf(":%v", listenPort))
  go messageFowarder(commReceive, receivedMsg, commSentStatus, commSend, sendCh)
  
  return commSentStatus, receivedMsg
}

func Send(sendCh chan c.CommData, targetAddr string, targetPort int, listenPort int) (<-chan c.ConnData, <-chan c.CommData) {
  commSend := make(chan c.CommData)
  commReceive := make(chan c.CommData, 1)
  commSentStatus := make(chan c.ConnData)
  receivedMsg := make(chan c.CommData)

  go listen(commReceive, fmt.Sprintf(":%v", listenPort), true)
  go broadcast(commSend, targetAddr, fmt.Sprintf(":%v", targetPort), fmt.Sprintf(":%v", listenPort))
  go messageFowarder(commReceive, receivedMsg, commSentStatus, commSend, sendCh)

  return commSentStatus, receivedMsg
}

func broadcast(send chan c.CommData, targetAddr string, targetPort string, listenPort string) {
  //local, err := net.ResolveUDPAddr("udp", listenPort)
  remoteAddr, _ := net.ResolveUDPAddr("udp", targetAddr + targetPort)
  conn, err := net.DialUDP("udp", nil, remoteAddr)
  defer conn.Close()

  if err != nil {
    panic(err)
  }
  
  log.Printf("server dial remote to %v", remoteAddr)

  for {
    msg := <- send
    convMsg, err := json.Marshal(msg)
    if err != nil {
      log.Printf("Convert json error: %v", err.Error())
    }
    if msg.AddrPort.Port > 0 {
      log.Printf("Sending back to %v", msg.AddrPort)
      conn.WriteToUDP(convMsg, &msg.AddrPort)
    } else {
      conn.Write(convMsg)
    }
  }
  
}

func listen(commReceive chan c.CommData, port string, readOnly bool) {
	listenAddr, _ := net.ResolveUDPAddr("udp", port)
	conn, err := net.ListenUDP("udp", listenAddr)
	defer conn.Close()

	if err != nil {
		panic(err)
	}
	
	log.Printf("server listening to %v", listenAddr)	

	var msg c.CommData
	for {
		buffer := make([]byte, 4096)
		length, replyAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
    if err := json.Unmarshal(buffer[:length], &msg); err != nil {
      panic(err)
    }
    msg.AddrPort = *replyAddr
    commReceive <- msg
	}
}

func messageFowarder(
  commReceive <-chan c.CommData, 
  receivedMsg chan<- c.CommData, 
  commSentStatus chan<- c.ConnData, 
  commSend chan<- c.CommData, 
  sendCh <-chan c.CommData) {
  for {
      select {
        case msg := <- commReceive:
          if msg.DataType == "Received"{
            response := c.ConnData{
              SenderIP: msg.SenderIP,
              MsgID: msg.MsgID,
              SendTime: time.Now(),
              Status: "Received",
            }
            receivedMsg <- msg
            commSentStatus <- response
          } else {
           response := c.CommData{
              Identifier: com_id,
              SenderIP: "127.0.0.1",
              ReceiverIP: msg.SenderIP,
              MsgID: msg.MsgID,
              AddrPort: msg.AddrPort,
              DataType: "Received",
              DataValue: time.Now(),
            }
            receivedMsg <- msg
            commSend <- response
          }
        case msg := <- sendCh: 
          commSend <- msg
          timeSent := c.ConnData{
            SenderIP: "127.0.0.1",
            MsgID: msg.MsgID,
            SendTime: time.Now(),
            Status: "Sent",
          }
          commSentStatus <- timeSent
      }
  }
}

func BuildMsg(receiverIP string, msgID string, dataType string, dataValue any) (*c.CommData) {
    return &c.CommData{
      Identifier: com_id,
      SenderIP: "hardcode_sender_ip",
      ReceiverIP: receiverIP,
      MsgID: msgID,
      DataType: dataType,
      DataValue: dataValue,
    }
}