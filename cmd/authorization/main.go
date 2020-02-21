package main

import (
	"fmt"
	"os"

	authorization "github.com/koverto/authorization/api"
	"github.com/koverto/authorization/internal/pkg/handler"
	"github.com/micro/go-micro/v2/config/source/env"

	"github.com/koverto/micro"
)

func main() {
	conf := &handler.Config{
		MongoUrl: "mongodb://localhost:27017",
	}

	service, err := micro.NewService("com.koverto.svc.authorization", conf, env.NewSource(env.WithStrippedPrefix("KOVERTO")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	h, err := handler.New(conf, service)
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
