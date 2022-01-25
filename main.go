package main

import (
	"flag"
	"fmt"
	fhandler "diction-car/handler/finder"
	chandler "diction-car/handler/maker"
	car "diction-carcar/pb/car"
	diction "grpccar/pb/diction"
	"log"
	"net"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	car.RegisterMakerServer(s, &chandler.Serviceserver{})
	diction.RegisterFinderServer(s, &fhandler.Serviceserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
