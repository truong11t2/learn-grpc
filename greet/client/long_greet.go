package main

import (
	"log"
	"time"

	pb "github.com/truong11t2/learn-go/greet/proto"
	"golang.org/x/net/context"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println(("doLongGreet was invoked"))

	reqs := []*pb.GreetRequest{
		{FirstName: "Truong"},
		{FirstName: "Test"},
		{FirstName: "John"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
