package main

import (
	"github.com/bolg-developers/BoLG-Server/api/bolg/v1/room"
	"github.com/bolg-developers/BoLG-Server/genproto/bolg/v1"
	"github.com/bolg-developers/BoLG-Server/server/config"
	"google.golang.org/grpc"
	"net"
)

func main() {
	port, err := net.Listen("tcp", config.Env().ServerPort)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	bolg.RegisterRoomServiceServer(srv, room.NewService(room.NewRepository(), room.NewPlayerRepository()))

	if err := srv.Serve(port); err != nil {
		panic(err)
	}
}
