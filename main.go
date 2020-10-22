package main

import (
	pb "github.com/amit/hello-service/rpc"
	"github.com/amit/hello-service/server"
	"net/http"
)

func main() {
	twirpHandler := pb.NewHelloWorldServer(&server.HelloWorldServer{}, nil)
	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8080", mux)
}
