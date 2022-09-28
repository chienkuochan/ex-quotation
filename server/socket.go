// auth: kunlun
// date: 2019-01-10
// description:
package server

import (
	"codec"
	"common"
	"encoding/json"
	"log"
	"net"
)

func Start() {
	address := "localhost:12345"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Fatalf("Analysis address error: %v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("Listener error: ", err)
		return
	}

	for {
		conn, err := listener.AcceptTCP()
		//conn.SetNoDelay(true)
		if err != nil {
			log.Fatalf("Accept error: %v", err)
			return
		}

		log.Printf("Client address is: %s", conn.RemoteAddr())

		go loopHandler(conn)
	}

}

func loopHandler(conn net.Conn) {
	//defer conn.Close()
	tmpBuffer := make([]byte, 0)
	//创建带缓冲的 chan
	ch := make(chan []byte, 16)
	go read(ch)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatalf("Read message error: %v", err)
			return
		}
		tmpBuffer = codec.Decoder(buffer[:n], ch)
		response := common.ResponseData{}
		_ = json.Unmarshal(tmpBuffer, &response)
		log.Printf("Analysis success: %v", response)
	}
}

func read(ch chan []byte) {
	select {
	case data := <-ch:
		log.Fatalf("Wtite data to chan: %v", data)
	}
}
