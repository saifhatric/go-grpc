package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/saif404/go-grpc/proto"
)

func CallSayHelloServer(client pb.GreetServiceClient, names * pb.NamesList){
	fmt.Printf("srever streaming started!\n");
	stream,err:= client.SayHelloServer(context.Background(),names)
	if err!=nil{
		log.Fatalf("could not send names : %v",names)
	}
	for {
		message,err:=stream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("error while streaming %v",err)
		}
		log.Println(message)
	}
	log.Printf("streaming finished!")
}