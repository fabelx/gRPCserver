package main

import (
	"fmt"
	api "github.com/fabelx/gRPCserver/pkg/api"
	"github.com/fabelx/gRPCserver/pkg/config"
	"github.com/fabelx/gRPCserver/pkg/divider"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	conf := config.GetConfig()
	port := fmt.Sprintf(":%s", conf.Port)
	s := grpc.NewServer()
	srv := &divider.Server{}
	api.RegisterDividerServer(s, srv)

	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
