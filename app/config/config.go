package config

import "github.com/kelseyhightower/envconfig"

type environment struct {
	ServerPort string `default:":50051"`
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
