package main

import (
	"context"
	"log"
	"time"

	pb "github.com/saif404/go-grpc/proto"
)

func CallSayHello(client pb.GreetServiceClient){
	ctx,cancel:= context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	
	res,err:=client.SayHello(ctx,&pb.NoParam{})
	if err!=nil{
		log.Fatalf("could not greet : %v\n", err)
	}
	log.Printf("%s",res.Message)
	
}