package main

import (
	"log"
	"testing"

	"github.com/jbactad/grpc-vs-rest/pb"
	"github.com/jbactad/grpc-vs-rest/testutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func BenchmarkGRPCHTTP2_4(b *testing.B) {
	const numWorkers = 4
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_8(b *testing.B) {
	const numWorkers = 8
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_16(b *testing.B) {
	const numWorkers = 16
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_32(b *testing.B) {
	const numWorkers = 32
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_64(b *testing.B) {
	const numWorkers = 64
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_128(b *testing.B) {
	const numWorkers = 128
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_256(b *testing.B) {
	const numWorkers = 256
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_512(b *testing.B) {
	const numWorkers = 512
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_1024(b *testing.B) {
	const numWorkers = 1024
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_2048(b *testing.B) {
	const numWorkers = 2048
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_4096(b *testing.B) {
	const numWorkers = 4096
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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

func BenchmarkGRPCHTTP2_8192(b *testing.B) {
	const numWorkers = 8192
	creds, err := credentials.NewClientTLSFromFile("./certs/server.crt", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("grpc:9093", grpc.WithTransportCredentials(creds))
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
