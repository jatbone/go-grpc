package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jatbone/go-grpc/proto"
)

func callSayHelloClientSteam(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with name: %v", name)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming closed")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}

	log.Printf("%v", res.Messages)

}
