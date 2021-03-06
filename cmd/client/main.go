package main

import (
	"context"
	"flag"
	"fmt"
	api "github.com/fabelx/gRPCserver/pkg/api"
	"github.com/fabelx/gRPCserver/pkg/config"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("Need two arguments (type: int).")
	}

	conf := config.GetConfig()
	port := fmt.Sprintf(":%s", conf.Port)

	x, err := strconv.Atoi(flag.Arg(0))

	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))

	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := api.NewDividerClient(conn)
	resp, err := client.Divide(context.Background(), &api.DivideRequest{X: int32(x), Y: int32(y)})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.String())
}
