package main

import (
	"io"
	"log"

	pb "github.com/truong11t2/learn-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg function was invoked")

	var res float32 = 0
	var i float32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			res /= i
			return stream.SendAndClose(&pb.AvgResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}
		log.Printf("Received number %d\n", req.Num)
		res += float32(req.Num)
		i++
	}
}
