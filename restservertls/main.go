package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jbactad/grpc-vs-rest/pb"
)

func handle(w http.ResponseWriter, req *http.Request) {
	log.Println("Request received.")
	random := pb.Random{RandomString: "a_random_string", RandomInt: 1984}

	bytes, err := json.Marshal(&random)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func main() {
	server := &http.Server{Addr: "localhost:9091", Handler: http.HandlerFunc(handle)}
	log.Println("Starting secured rest server at port")
	log.Fatal(server.ListenAndServeTLS("../certs/server.crt", "../certs/server.key"))
}
