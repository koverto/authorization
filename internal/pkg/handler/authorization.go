package handler

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"

	"github.com/dgrijalva/jwt-go/v4"
	authz "github.com/koverto/authorization/api"
	"github.com/koverto/authorization/pkg/claims"

	"github.com/koverto/errors"
	"github.com/koverto/micro"
	"github.com/koverto/mongo"
)

type Authorization struct {
	*Config
	*micro.Service
	client     mongo.Client
	privateKey *ecdsa.PrivateKey
}

type Config struct {
	MongoUrl   string `json:"mongourl"`
	PrivateKey string `json:"privatekey"`
}

func New(conf *Config, service *micro.Service) (*Authorization, error) {
	client, err := mongo.NewClient(conf.MongoUrl, service.Name)
	if err != nil {
		return nil, err
	}

	if err := client.Connect(); err != nil {
		return nil, err
	}

	pkey, err := x509.ParseECPrivateKey([]byte(conf.PrivateKey))
	if err != nil {
		return nil, err
	}

	return &Authorization{conf, service, client, pkey}, nil
}

func (a *Authorization) Create(ctx context.Context, in *authz.TokenRequest, out *authz.Token) (err error) {
	c := claims.New(in.GetUserID())
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	out.Token, err = token.SignedString(a.privateKey)
	return err
}

func (a *Authorization) Validate(ctx context.Context, in *authz.Token, out *authz.TokenResponse) error {
	return errors.NotImplemented(a.ID)
}

func (a *Authorization) Invalidate(ctx context.Context, in *authz.Token, out *authz.TokenResponse) error {
	return errors.NotImplemented(a.ID)
}
