package test

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	startServer()
	startClient()
	exitCode := m.Run()
	stopServer()
	closeClient()
	os.Exit(exitCode)
}

func TestHello(t *testing.T) {
	ctx := context.Background()
	log.Println("Testing Hello Route")
	for _, request := range requestList {
		var reply string
		err := clientConn.Invoke(ctx, "/twirp/helloservice.HelloWorld/Hello", &request, &reply)
		log.Println(reply)
		if err != nil {
			log.Println(err)
			t.Error(err)
		}
	}
}
