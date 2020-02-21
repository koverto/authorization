package handler

import (
	"context"

	authz "github.com/koverto/authorization/api"

	"github.com/koverto/errors"
	"github.com/koverto/micro"
	"github.com/koverto/mongo"
)

type Authorization struct {
	*Config
	*micro.Service
	client mongo.Client
}

type Config struct {
	MongoUrl string `json:"mongourl"`
}

func New(conf *Config, service *micro.Service) (*Authorization, error) {
	client, err := mongo.NewClient(conf.MongoUrl, service.Name)
	if err != nil {
		return nil, err
	}

	if err := client.Connect(); err != nil {
		return nil, err
	}

	return &Authorization{conf, service, client}, nil
}

func (a *Authorization) Create(ctx context.Context, in *authz.TokenClaims, out *authz.Token) error {
	return errors.NotImplemented(a.ID)
}

func (a *Authorization) Validate(ctx context.Context, in *authz.Token, out *authz.TokenClaims) error {
	return errors.NotImplemented(a.ID)
}

func (a *Authorization) Invalidate(ctx context.Context, in *authz.TokenClaims, out *authz.TokenClaims) error {
	return errors.NotImplemented(a.ID)
}
