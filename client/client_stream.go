package main

import (
	"context"
	"log"
	"time"

	pb "github.com/saif404/go-grpc/proto"
)

func CallSayHelloClient(client pb.GreetServiceClient,names *pb.NamesList){
	log.Printf("Streaming Started!")
	stream,err:= client.SayHelloClient(context.Background())
	if err!=nil{
		log.Fatalf("could not send the names %v ",err)
	}
	for _,name:= range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err:=stream.Send(req);err!=nil{
			log.Fatalf("Error while sending %v",req)
		}
		log.Printf("Sent the request with name %s",name)
		time.Sleep(2 *time.Second)
	}
	res,err:= stream.CloseAndRecv()
	log.Printf("Client Streaming finished!")
	if err!=nil{
		log.Fatalf("error while resiving %v",err)
	}
	log.Printf("%v",res.Messages)
}