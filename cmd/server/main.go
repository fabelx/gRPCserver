package main

import (
	api "gRPCserver/pkg/api"
	"gRPCserver/pkg/divider"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &divider.Server{}
	api.RegisterDividerServer(s, srv)

	l, err := net.Listen("tcp", ":8191")

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
