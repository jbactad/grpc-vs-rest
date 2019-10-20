package main

import (
	"encoding/json"
	"github.com/jbactad/grpc-vs-rest/logger"
	"go.uber.org/zap"
	"log"
	"net/http"

	"github.com/jbactad/grpc-vs-rest/pb"
)

func handle(w http.ResponseWriter, req *http.Request) {
	logger.Logger.Info("Request received.")
	logger.Logger.Info("headers intercepted", zap.Any("header", req.Header))
	random := pb.Random{RandomString: "a_random_string", RandomInt: 1984}

	bytes, err := json.Marshal(&random)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func main() {
	server := &http.Server{Addr: "0.0.0.0:9091", Handler: http.HandlerFunc(handle)}
	logger.Logger.Info("Starting secured rest server at port 9091")
	log.Fatal(server.ListenAndServeTLS("/etc/certs/server.crt", "/etc/certs/server.key"))
}
