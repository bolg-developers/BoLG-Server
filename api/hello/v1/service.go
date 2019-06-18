package v1

import (
	"context"
	pbhellov1 "github.com/bolg-developers/BoLG-Server/genproto/hello/v1"
	_type "github.com/bolg-developers/BoLG-Server/genproto/type"
	"log"
)

type HelloService struct{}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (HelloService) Greet(context.Context, *_type.Empty) (*pbhellov1.GreetResponse, error) {
	log.Println("call HelloService.Greet")
	return &pbhellov1.GreetResponse{Message: "hello!!"}, nil
}
