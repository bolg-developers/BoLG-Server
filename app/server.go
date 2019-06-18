package main

import (
	hellov1 "github.com/bolg-developers/BoLG-Server/api/hello/v1"
	"github.com/bolg-developers/BoLG-Server/app/config"
	pbhellov1 "github.com/bolg-developers/BoLG-Server/genproto/hello/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port, err := net.Listen("tcp", config.Env().ServerPort)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pbhellov1.RegisterHelloServer(srv, hellov1.NewHelloService())

	log.Println("running!")
	if err := srv.Serve(port); err != nil {
		panic(err)
	}
}
