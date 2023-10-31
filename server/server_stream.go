package main

import (
	"log"
	"time"

	pb "github.com/saif404/go-grpc/proto"
)

func (s *helloServer) SayHelloServer(req *pb.NamesList, stream pb.GreetService_SayHelloServerServer )error{
	log.Printf("got request with names : %v",req.Names)
	for _,name := range req.Names{
		res := &pb.HelloResponse{
			Message: "Hello_" + name,
		}
		if err:= stream.Send(res);err!=nil{
			return err
		}
		time.Sleep(time.Second *2)
	}
	return nil
}