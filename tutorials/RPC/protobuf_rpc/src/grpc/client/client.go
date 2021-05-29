package main 

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"golang.org/x/net/context"

	. "grpc/hello"
)

func main(){
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())

	if err != nil{
		log.Fatal("dialing: ", err)
	}

	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "Hello from client"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}