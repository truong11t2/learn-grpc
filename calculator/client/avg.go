package main

import (
	"context"
	"log"
	"time"

	pb "github.com/truong11t2/learn-go/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("doAvg was invoked")

	reqs := []*pb.AvgRequest{
		{Num: 1},
		{Num: 2},
		{Num: 7},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)
}
