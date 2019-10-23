#!/usr/bin/env bash

go test -run=^$ -bench=RestHTTP2 -benchtime=1000x > rest.out &&
go test -run=^$ -bench=GRPCHTTP2 -benchtime=1000x > grpc.out &&
cat rest.out grpc.out| benchgraph -title="gRPC vs Rest"

