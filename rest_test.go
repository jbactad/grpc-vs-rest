package main

import (
	"net/http"
	"testing"

	"github.com/jbactad/grpc-vs-rest/pb"
	"github.com/jbactad/grpc-vs-rest/testutils"
)

func BenchmarkRestHTTP11_4(b *testing.B) {
	const numWorkers = 4
	client := http.Client{}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path: "http://localhost:9090",
			Random: &pb.Random{
				RandomInt:    2019,
				RandomString: "a_string",
			},
		}
	}
}

func BenchmarkRestHTTP2_4(b *testing.B) {
	const numWorkers = 4
	client := http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path: "https://localhost:9091",
			Random: &pb.Random{
				RandomInt:    2019,
				RandomString: "a_string",
			},
		}
	}
}

func BenchmarkRestHTTP2_8(b *testing.B) {
	const numWorkers = 8 
	client := http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path: "https://localhost:9091",
			Random: &pb.Random{
				RandomInt:    2019,
				RandomString: "a_string",
			},
		}
	}
}

func BenchmarkRestHTTP2_16(b *testing.B) {
	const numWorkers = 16 
	client := http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path: "https://localhost:9091",
			Random: &pb.Random{
				RandomInt:    2019,
				RandomString: "a_string",
			},
		}
	}
}

func BenchmarkRestHTTP2_32(b *testing.B) {
	const numWorkers = 32 
	client := http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path: "https://localhost:9091",
			Random: &pb.Random{
				RandomInt:    2019,
				RandomString: "a_string",
			},
		}
	}
}
