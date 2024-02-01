package main

import (
	"log"
	"net"

	pb "github.com/truong11t2/learn-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	opts := []grpc.ServerOption{}

	tls := true // change to false if needed

	if tls {
		certFile := "../ssl/server.crt"
		keyFile := "../ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)

		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	//s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
}
