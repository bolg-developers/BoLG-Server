package config

import "github.com/kelseyhightower/envconfig"

const (
	RoomIdClaimsKey = "roomId"
	IrIdClaimsKey   = "IrId"
)

type environment struct {
	ServerPort     string `default:":50051"`
	MaxRoomCount   int    `default:"16"`
	MaxIrId        int    `default:"15"`
	TokenSecretKey string `required:"true"`
}

var env environment

func Env() *environment {
	return &env
}

func init() {
	err := envconfig.Process("BOLG", &env)
	if err != nil {
		panic(err)
	}
}
