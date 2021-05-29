package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"golang.org/x/net/context"

	. "grpc/hello"
)

type HelloServiceImpl struct{}

//kiểu String được định nghĩa trong protobuf
func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error){
	//hàm .GetValue() đã được tạo ra trong file hello.pb.go
	reply := &String{Value: "Hello: " + args.GetValue()}
	return reply, nil
}

func main(){
	//khởi tạo một đối tượng gRPC service
	grpcServer := grpc.NewServer()

	//đăng ký service với grpcServer (của gRPC plugin)
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	//cung cấp gRPC service trên port "1234"
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}