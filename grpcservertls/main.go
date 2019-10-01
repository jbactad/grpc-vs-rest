package main

import (
	"context"
	"log"
	"net"

	"github.com/jbactad/grpc-vs-rest/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func main() {
	creds, err := credentials.NewServerTLSFromFile("../certs/server.crt", "../certs/server.key")
	if err != nil {
		log.Fatalf("Failed to setup TLS: %v", err)
	}

	lis, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterRandomServiceServer(s, &server{})

	log.Println("Starting secured gRPC server at port 9093")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) DoSomething(_ context.Context, random *pb.Random) (*pb.Random, error) {
	log.Println("Request received.")
	random.RandomString = "[Updated] " + random.RandomString
	return random, nil
}
