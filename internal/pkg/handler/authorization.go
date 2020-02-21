package handler

import (
	"context"

	authorization "github.com/koverto/authorization/api"

	"github.com/koverto/errors"
	"github.com/koverto/mongo"
)

type Authorization struct {
	*Config
	client mongo.Client
}

func New(conf *Config) (*Authorization, error) {
	client, err := mongo.NewClient(conf.MongoUrl, conf.Name)
	if err != nil {
		return nil, err
	}

	if err := client.Connect(); err != nil {
		return nil, err
	}

	return &Authorization{conf, client}, nil
}

func (a *Authorization) Create(ctx context.Context, in *authorization.TokenClaims, out *authorization.Token) error {
	return errors.NotImplemented(a.ID())
}

func (a *Authorization) Validate(ctx context.Context, in *authorization.Token, out *authorization.TokenClaims) error {
	return errors.NotImplemented(a.ID())
}

func (a *Authorization) Invalidate(ctx context.Context, in *authorization.TokenClaims, out *authorization.TokenClaims) error {
	return errors.NotImplemented(a.ID())
}
