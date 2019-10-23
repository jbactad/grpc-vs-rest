package main

import (
	"net/http"
	"testing"

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
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
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
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
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
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
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
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
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
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
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
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
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
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
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
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
		}
	}
}

func BenchmarkRestHTTP2_1024(b *testing.B) {
	const numWorkers = 1024
	client := http.Client{}
	client.Transport = &http2.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
		}
	}
}

func BenchmarkRestHTTP2_2048(b *testing.B) {
	const numWorkers = 2048
	client := http.Client{}
	client.Transport = &http2.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
		}
	}
}

func BenchmarkRestHTTP2_4096(b *testing.B) {
	const numWorkers = 4096
	client := http.Client{}
	client.Transport = &http2.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
		}
	}
}

func BenchmarkRestHTTP2_8192(b *testing.B) {
	const numWorkers = 8192
	client := http.Client{}
	client.Transport = &http2.Transport{
		TLSClientConfig: testutils.CreateTLSConfigWithCustomCert("./certs/server.crt"),
	}
	requestQueue := make(chan testutils.Request)
	defer testutils.StartWorkers(&requestQueue, numWorkers, testutils.StartPostWorkerFunc(client))()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestQueue <- testutils.Request{
			Path:        "https://grpc:9091",
			Transaction: *testutils.SampleTransaction(),
		}
	}
}
