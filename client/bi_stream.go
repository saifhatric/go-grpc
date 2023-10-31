package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/saif404/go-grpc/proto"
)
func CallSayHelloBiDirectional(client pb.GreetServiceClient,names *pb.NamesList){
	log.Printf("Bi_Directional Streaming started!")
	stream,err:= client.SayHelloBiDirectional(context.Background())
	if err!=nil{
		log.Fatalf("could not send names to the server %v",err)
	}
	waitc := make(chan struct{})
	go func(){
		for{
			message,err:= stream.Recv()
			if err == io.EOF{
				break
			}
			if err!=nil{
				log.Fatalf("Error happend %v",err)
			}
			log.Println(message)
		}
		close(waitc)
	}()
	for _,name := range names.Names{
		req:= &pb.HelloRequest{
			Name:name,
		}
		if err:=stream.Send(req);err!=nil{
			log.Fatalf("Error while sending the request %v",err)
		}
		time.Sleep(time.Second *2)

	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bi Directional streaming finished!")
}