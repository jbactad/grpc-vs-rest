package main

import (
	"log"
	"net"

	"github.com/jbactad/grpc-vs-rest/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRandomServiceServer(s, &server{})

	log.Println("Starting gRPC server at pot 9092")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) DoSomething(_ context.Context, random *pb.Random) (*pb.Random, error) {
	log.Println("Request received.")
	random.RandomString = "[Updated] " + random.RandomString
	return random, nil
}
