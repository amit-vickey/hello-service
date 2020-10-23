package server

import (
	"bytes"
	"encoding/json"
	pb "github.com/amit/hello-service/rpc"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	requestList = [] pb.HelloReq{
		{
			Subject: "sample",
		},
		{
			Subject: "amit",
		},
	}
)

func Mux() *http.ServeMux {
	twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{}, nil)
	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	return mux
}


func TestServer(t *testing.T) {
	log.Println("Testing Hello Route")
	for _, request := range requestList {
		requestByte, _ := json.Marshal(request)
		httpRequest, _ := http.NewRequest("POST", "http://127.0.0.1:8080/twirp/helloservice.HelloWorld/Hello", bytes.NewReader(requestByte))
		httpRequest.Header.Add("content-type", "application/json")
		rec := httptest.NewRecorder()
		Mux().ServeHTTP(rec, httpRequest)
		response := rec.Result()
		defer response.Body.Close()
		log.Println("Response status : ", response.Status)
		log.Println("Response status code : ", response.StatusCode)
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("Error reading response : %v", err)
		}
		log.Println("Response body : ", string(responseBody))
	}
}
