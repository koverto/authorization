package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	authorization "github.com/koverto/authorization/api"
	"github.com/koverto/authorization/internal/pkg/handler"

	"github.com/koverto/micro/v2"
	"github.com/micro/go-micro/v2/config/source/env"
)

func main() {
	pkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	enc, err := x509.MarshalECPrivateKey(pkey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf := &handler.Config{
		MongoUrl: "mongodb://localhost:27017",
		PrivateKey: string(pem.EncodeToMemory(&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: enc,
		})),
	}

	service, err := micro.NewService(authorization.Name, conf, env.NewSource(env.WithStrippedPrefix("KOVERTO")))
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
