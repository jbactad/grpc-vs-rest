FROM golang:1.13.2-stretch as build

ENV WORKDIR=/go/src/github.com/jbactad/grpc-vs-rest \
    CGO_ENABLED=0

RUN mkdir -p $WORKDIR

WORKDIR $WORKDIR

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/grpcservertls grpcservertls/main.go \
    && go build -o ./bin/restservertls restservertls/main.go

FROM alpine

COPY --from=build /go/src/github.com/jbactad/grpc-vs-rest/bin /bin/

CMD /bin/grpcservertls
