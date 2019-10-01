package main

import (
	"context"
	"log"
	"testing"

	"github.com/jbactad/grpc-vs-rest/pb"
	"github.com/jbactad/grpc-vs-rest/testutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func BenchmarkGRPC_WithWokers(b *testing.B) {
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	client := pb.NewRandomServiceClient(conn)
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartGRPCWorkerFunc(client))()
	b.ResetTimer() // don't count worker initialization time

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Random: &pb.Random{
				RandomInt:    2019,
				RandomString: "a_string",
			},
		}
	}
}

func BenchmarkGRPC_NoWorkers(b *testing.B) {
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	client := pb.NewRandomServiceClient(conn)

	request := testutils.Request{
		Random: &pb.Random{
			RandomInt:    2019,
			RandomString: "a_string",
		},
	}

	_, err = client.DoSomething(context.TODO(), request.Random)
	if err != nil {
		log.Printf("Error sending grpc request: %v\n", err)
	}
}

func BenchmarkTlsGRPC_WithWokers(b *testing.B) {
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "localhost")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("localhost:9093", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	client := pb.NewRandomServiceClient(conn)
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartGRPCWorkerFunc(client))()
	b.ResetTimer() // don't count worker initialization time

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Random: &pb.Random{
				RandomInt:    2019,
				RandomString: "a_string",
			},
		}
	}
}

func BenchmarkTlsGRPC_NoWorkers(b *testing.B) {
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "localhost")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("localhost:9093", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	client := pb.NewRandomServiceClient(conn)

	request := testutils.Request{
		Random: &pb.Random{
			RandomInt:    2019,
			RandomString: "a_string",
		},
	}

	_, err = client.DoSomething(context.TODO(), request.Random)
	if err != nil {
		log.Printf("Error sending grpc request: %v\n", err)
	}
}
