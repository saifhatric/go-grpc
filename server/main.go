package main

import (
	"log"
	"net"

	pb "github.com/saif404/go-grpc/proto"
	"google.golang.org/grpc"
)
const(
	port=":8080"
)
type helloServer struct{
	pb.GreetServiceServer
}
func main(){
	listen,err := net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("Failed to start the server %v\n",err)
	}
	grpcServer:= grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer,&helloServer{})
	log.Printf("server started at %v\n",listen.Addr())
	if err:= grpcServer.Serve(listen);err!=nil{
		log.Fatalf("Failed to start %v\n",err)
	}

}