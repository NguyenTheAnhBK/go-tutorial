package main 

import (
	"fmt"
	"log"
	"net/rpc"
)

func main(){
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil{
		log.Fatal("dialing: ", err)
	}

	//biến chứa giá trị trả về sau lời gọi rpc
	var reply string

	err = client.Call("HelloService.Hello", "Hello Server", &reply)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(reply)
}