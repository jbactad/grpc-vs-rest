package main

import (
	"net/http"
	"testing"

	"github.com/jbactad/grpc-vs-rest/pb"
	"github.com/jbactad/grpc-vs-rest/testutils"
)

func BenchmarkRest_WithWorker(b *testing.B) {
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

func BenchmarkRest_NoWorker(b *testing.B) {
	client := http.Client{}

	request := testutils.Request{
		Path: "http://localhost:9090",
		Random: &pb.Random{
			RandomInt:    2019,
			RandomString: "a_string",
		},
	}

	testutils.Get(client, "http://localhost:9090", request.Random)
}

func BenchmarkTlsRest_WithWorker(b *testing.B) {
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

func BenchmarkTlsRest_NoWorker(b *testing.B) {
	client := http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}

	request := testutils.Request{
		Path: "https://localhost:9091",
		Random: &pb.Random{
			RandomInt:    2019,
			RandomString: "a_string",
		},
	}

	testutils.Get(client, "https://localhost:9091", request.Random)
}
