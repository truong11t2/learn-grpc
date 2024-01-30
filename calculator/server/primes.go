package main

import (
	"log"

	pb "github.com/truong11t2/learn-go/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)
	num := in.Num
	var i int32

	for i = 2; i < num; i++ {
		for num%i == 0 {
			// Send i to client
			stream.Send(&pb.PrimeResponse{
				Prime: i,
			})
			num = num / i
		}
	}

	return nil
}
