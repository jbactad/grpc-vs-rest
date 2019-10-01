package testutils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/jbactad/grpc-vs-rest/pb"
)

const stopRequestPath = "STOP"

type WorkerFunc func(*chan Request, *sync.WaitGroup)

func CreateTLSConfigWithCustomCert(certPath string) *tls.Config {
	caCert, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}

	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(caCert)
	if !ok {
		log.Fatal("Error appending certificate to pool")
	}

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	return tlsConfig
}

func StartGRPCWorkerFunc(client pb.RandomServiceClient) WorkerFunc {
	return func(requestQueue *chan Request, wg *sync.WaitGroup) {
		go func() {
			for {
				request := <-*requestQueue
				if request.Path == stopRequestPath {
					wg.Done()
					return
				}
				_, err := client.DoSomething(context.TODO(), request.Random)
				if err != nil {
					log.Printf("Error sending grpc request: %v\n", err)
				}
			}
		}()
	}
}

func StartWorkers(requestQueue *chan Request, numWorkers int, startWorkerFunc WorkerFunc) func() {
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		startWorkerFunc(requestQueue, &wg)
	}
	return func() {
		wg.Add(numWorkers)
		stopRequest := Request{Path: stopRequestPath}
		for i := 0; i < numWorkers; i++ {
			*requestQueue <- stopRequest
		}
		wg.Wait()
	}
}

func StartWorkerFunc(client http.Client) WorkerFunc {
	return func(requestQueue *chan Request, wg *sync.WaitGroup) {
		go func() {
			for {
				request := <-*requestQueue
				if request.Path == stopRequestPath {
					wg.Done()
					return
				}
				Get(client, request.Path, request.Random)
			}
		}()
	}
}

func StartPostWorkerFunc(client http.Client) WorkerFunc {
	return func(requestQueue *chan Request, wg *sync.WaitGroup) {
		go func() {
			for {
				request := <-*requestQueue
				if request.Path == stopRequestPath {
					wg.Done()
					return
				}
				Post(client, request.Path, request.Random, request.Random)
			}
		}()
	}
}
