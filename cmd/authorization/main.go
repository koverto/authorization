package main

import (
	"fmt"
	"os"

	authorization "github.com/koverto/authorization/api"
	"github.com/koverto/authorization/internal/pkg/handler"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/source/env"
)

func main() {
	service := micro.NewService(micro.Name("authorization"))
	service.Init()

	conf, err := handler.NewConfig("authorization", env.NewSource(env.WithStrippedPrefix("KOVERTO")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	h, err := handler.New(conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := authorization.RegisterAuthorizationHandler(service.Server(), h); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
