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
	server := &http.Server{Addr: "localhost:9090", Handler: http.HandlerFunc(handle)}
	log.Println("Starting rest server at port 9090")
	log.Fatal(server.ListenAndServe())
}
