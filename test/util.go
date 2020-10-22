package test

import (
	"context"
	pb "github.com/amit/hello-service/rpc"
	"github.com/amit/hello-service/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"net/http"
)

const bufSize = 1024 * 1024

var (
	grpcServer *grpc.Server
	//httpServer http.Server
	listener *bufconn.Listener
	clientConn *grpc.ClientConn
)

func serveGRPC(l net.Listener) {
	grpcServer := grpc.NewServer()
	log.Println("Starting GRPC Server")
	if err := grpcServer.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}

func serveHTTP1(l net.Listener) {
	twirpHandler := pb.NewHelloWorldServer(&server.HelloWorldServer{}, nil)
	s := &http.Server{
		Handler: twirpHandler,
	}
	log.Println("Starting HTTP Server")
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}

func startServer() {
	listener = bufconn.Listen(bufSize)
	/*
	CMUX Use Link: https://godoc.org/github.com/soheilhy/cmux#example
	mux := cmux.New(listener)
	grpcL := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := mux.Match(cmux.Any())
	go serveGRPC(grpcL)
	go serveHTTP1(httpL)
	go func() {
		log.Println("Starting Mux Server")
		if err := mux.Serve(); err != nil {
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Println("true")
			} else {
				log.Println("false")
			}
			log.Fatalln(err)
		}
	}()*/
	/*mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)*/
	//TODO: Need a way to attach our server to a grpc server
	grpcServer = grpc.NewServer()
	log.Println(grpcServer.GetServiceInfo())
	// ServiceDesc example: /Users/amitvickey/go/src/cloud.google.com/go/pubsub/internal/benchwrapper/proto/pubsub.pb.go:190
	serviceDescriptor := grpc.ServiceDesc{
		ServiceName: "payout-links",
		HandlerType: (*server.HelloWorldServer)(nil),
		Methods: []grpc.MethodDesc {
			{
				MethodName: "Hello",
				Handler: _Hello_RPC_Handler,
			},
		},
	}
	grpcServer.RegisterService(&serviceDescriptor, nil)
	log.Println(grpcServer.GetServiceInfo())
	/*httpServer = http.Server{
		Handler: mux,
	}
	go func() {
		log.Println("Starting grpc server")
		if err := grpcServer.Serve(listener); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("Server exited with error: %v", err)
			}
		}
	}()*/
}

func _Hello_RPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	worldServer := srv.(server.HelloWorldServer)
	return worldServer.Hello(ctx, in)
}

func startClient() {
	log.Println("Starting test client")
	var err error
	ctx := context.Background()
	clientConn, err = grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(bufDialer))
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func stopServer() {
	log.Println("Stopping grpc server")
	//httpServer.Shutdown(context.Background())

	grpcServer.Stop()
}

func closeClient() {
	log.Println("Stopping test client")
	err := clientConn.Close()
	if err != nil {
		log.Fatalf("Error stopping client : %v", err)
	}
}
