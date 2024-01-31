package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/truong11t2/learn-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Printf("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Num: -999},
		{Num: 23},
		{Num: 890},
		{Num: 70},
		{Num: -1},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Println("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Max number: %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
