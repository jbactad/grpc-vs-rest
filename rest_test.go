package main

import (
	"net/http"
	"testing"

	"github.com/jbactad/grpc-vs-rest/pb"
	"github.com/jbactad/grpc-vs-rest/testutils"

	"golang.org/x/net/http2"
)

func BenchmarkRestHTTP2_4(b *testing.B) {
	const numWorkers = 4
	client := http.Client{}
	client.Transport = &http2.Transport{
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
	client.Transport = &http2.Transport{
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
	client.Transport = &http2.Transport{
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
	client.Transport = &http2.Transport{
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

func BenchmarkRestHTTP2_64(b *testing.B) {
	const numWorkers = 64
	client := http.Client{}
	client.Transport = &http2.Transport{
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

func BenchmarkRestHTTP2_128(b *testing.B) {
	const numWorkers = 128
	client := http.Client{}
	client.Transport = &http2.Transport{
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

func BenchmarkRestHTTP2_256(b *testing.B) {
	const numWorkers = 256
	client := http.Client{}
	client.Transport = &http2.Transport{
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

func BenchmarkRestHTTP2_512(b *testing.B) {
	const numWorkers = 512
	client := http.Client{}
	client.Transport = &http2.Transport{
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
