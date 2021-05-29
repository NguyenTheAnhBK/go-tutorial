package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error{
	*reply = "[Server] reply: " + request
	return nil
}

func main(){
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")

	if err != nil{
		log.Fatal("Listen TCP error: ", err)
	}

	for{
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		//phục vụ client trên một gorountine khác
		//để giải phóng main thread tiếp tục vòng lặp
		go rpc.ServeConn(conn)
	}
}