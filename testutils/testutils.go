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

func StartGRPCWorkerFunc(client pb.TransactionServiceClient) WorkerFunc {
	return func(requestQueue *chan Request, wg *sync.WaitGroup) {
		go func() {
			for {
				request := <-*requestQueue
				if request.Path == stopRequestPath {
					wg.Done()
					return
				}
				_, err := client.CreateTransaction(context.TODO(), SampleTransactionRequest())
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
				resp := &pb.Transaction{}

				err := Get(client, request.Path, resp)
				if err != nil {
					log.Fatal(err)
				}
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
				resp := &pb.Transaction{}
				err := Post(client, request.Path, request.Transaction, resp)
				if err != nil {
					log.Fatal(err)
				}
			}
		}()
	}
}

func SampleTransactionRequest() *pb.CreateTransactionRequest {
	return &pb.CreateTransactionRequest{
		Transaction: SampleTransaction(),
	}
}
func SampleTransaction() *pb.Transaction {
	return &pb.Transaction{
		Date:                 0,
		ExternalId:           0,
		Source:               0,
		Type:                 0,
		PropertyReference:    "",
		Price:                0,
		VerificationStatus:   "",
		RejectionReason:      "",
		VerifiedAt:           0,
		Furnishing:           0,
		DeveloperName:        "",
		ChequeCount:          0,
		BedroomCount:         0,
		BathroomCount:        0,
		BuiltUpArea:          0,
		PlotSize:             0,
		UnitNumber:           "",
		FloorLevel:           "",
		Location:             nil,
		PropertyType:         "",
		Seller:               nil,
		Buyer:                nil,
		Landlord:             nil,
		Tenant:               nil,
		IsChillerInclusive:   false,
		HasStudyRoom:         false,
		HasMaidRoom:          false,
		Views:                nil,
		ProjectStatus:        "",
		ContractDuration:     "",
		Commission:           nil,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
}
