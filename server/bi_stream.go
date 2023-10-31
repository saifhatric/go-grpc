package main

import (
	"io"
	"log"

	pb "github.com/saif404/go-grpc/proto"
)

func (s *helloServer) SayHelloBiDirectional(stream pb.GreetService_SayHelloBiDirectionalServer)error{
	for{
		req,err:= stream.Recv()
		if err== io.EOF{
			return nil
		}
		if err!=nil{
			return err
		}
		log.Printf("Got request with name %v",req.Name)
		res := &pb.HelloResponse{
			Message: "Hello_"+req.Name,
		}
		if err:=stream.Send(res);err!=nil{
			return err
		}
	}
}