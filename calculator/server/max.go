package main

import (
	"io"
	"log"

	pb "github.com/truong11t2/learn-go/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function was invoked with \n")

	var res int32 = -99999

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if res < req.Num {
			res = req.Num
		}

		err = stream.Send(&pb.MaxResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
