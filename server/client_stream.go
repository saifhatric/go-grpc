package main

import (
	"io"
	"log"

	pb "github.com/saif404/go-grpc/proto"
)

func (s *helloServer) SayHelloClient(stream pb.GreetService_SayHelloClientServer)error{
	var messages []string

	for{
		req,err:=stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err!=nil{
			return err
		}
		log.Printf("Got a request with name : %v",req.Name)
		messages = append(messages, "Hello_",req.Name)
	}
	
}