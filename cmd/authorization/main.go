package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"

	authorization "github.com/koverto/authorization/api"
	"github.com/koverto/authorization/internal/pkg/handler"

	"github.com/koverto/micro/v2"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/rs/zerolog/log"
)

func main() {
	pkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("could not generate signing key")
	}

	enc, err := x509.MarshalECPrivateKey(pkey)
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("could not read private key")
	}

	conf := &handler.Config{
		MongoURL: "mongodb://localhost:27017",
		PrivateKey: string(pem.EncodeToMemory(&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: enc,
		})),
	}

	service, err := micro.NewService(authorization.Name, conf, env.NewSource(env.WithStrippedPrefix("KOVERTO")))
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("could not initialize service")
	}

	h, err := handler.New(conf, service)
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("could not build handler")
	}

	if err := authorization.RegisterAuthorizationHandler(service.Server(), h); err != nil {
		log.Fatal().AnErr("error", err).Msg("could not register handler with service")
	}

	if err := service.Run(); err != nil {
		log.Fatal().AnErr("error", err).Msg("error running service")
	}
}
