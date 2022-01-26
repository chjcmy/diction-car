package main

import (
	fhandler "diction-car/handler/finder"
	mhandler "diction-car/handler/maker"
	diction "diction-car/pb/diction"
	parallel "diction-car/pb/parallel"
	"flag"
	"fmt"
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
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	parallel.RegisterParallelServer(s, &mhandler.Serviceserver{})
	diction.RegisterFinderServer(s, &fhandler.Serviceserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
