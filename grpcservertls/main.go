package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/jbactad/grpc-vs-rest/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/jbactad/grpc-vs-rest/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (s *server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	logger.Logger.Info("Request received.")
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		logger.Logger.Info("headers intercepted", zap.Any("header", md))
	}

	return &pb.CreateTransactionResponse{
		Transaction: req.Transaction,
	}, nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("/etc/certs/server.crt", "/etc/certs/server.key")
	if err != nil {
		log.Fatalf("Failed to setup TLS: %v", err)
	}

	lis, err := net.Listen("tcp", "0.0.0.0:9093")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger.Logger),
		)),
	}

	s := grpc.NewServer(opts...)

	reflection.Register(s)

	// pb.RegisterRandomServiceServer(s, &server{})
	pb.RegisterTransactionServiceServer(s, &server{})

	logger.Logger.Info("Starting secured gRPC server at port 9093")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
