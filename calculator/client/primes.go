package main

import (
	"context"
	"io"
	"log"

	pb "github.com/truong11t2/learn-go/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{
		Num: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to connect: %v\n", err)
		}

		log.Printf("Primes: %d\n", msg.Prime)
	}
}
