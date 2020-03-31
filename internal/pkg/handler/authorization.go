package handler

import (
	"context"
	"crypto/ecdsa"
	"time"

	authz "github.com/koverto/authorization/api"
	"github.com/koverto/authorization/pkg/claims"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/koverto/micro/v2"
	"github.com/koverto/mongo"
	"github.com/koverto/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mmongo "go.mongodb.org/mongo-driver/mongo"
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

type invalidToken struct {
	TokenID   *uuid.UUID `bson:"_id"`
	ExpiresAt *time.Time
}

func New(conf *Config, service *micro.Service) (*Authorization, error) {
	client, err := mongo.NewClient(conf.MongoUrl, "authorization")
	if err != nil {
		return nil, err
	}

	if err := client.Connect(); err != nil {
		return nil, err
	}

	pkey, err := jwt.ParseECPrivateKeyFromPEM([]byte(conf.PrivateKey))
	if err != nil {
		return nil, err
	}

	return &Authorization{conf, service, client, pkey}, nil
}

func (a *Authorization) Create(ctx context.Context, in *authz.Claims, out *authz.Token) (err error) {
	c := claims.New(in.GetSubject())
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	out.Token, err = token.SignedString(a.privateKey)
	return err
}

func (a *Authorization) Validate(ctx context.Context, in *authz.Token, out *authz.Claims) error {
	token, claims, err := a.parseToken(in)
	if err != nil {
		return err
	}

	invalidated, err := a.tokenInvalidated(ctx, claims.ID)

	if token.Valid && !invalidated && err == nil {
		out.ID = claims.ID
		out.Subject = claims.Subject
		out.ExpiresAt = &claims.ExpiresAt.Time
	}

	return err
}

func (a *Authorization) Invalidate(ctx context.Context, in *authz.Claims, out *authz.Claims) error {
	if invalidated, err := a.tokenInvalidated(ctx, in.GetID()); invalidated {
		return nil
	} else if err != nil {
		return err
	}

	invalid := &invalidToken{
		TokenID:   in.GetID(),
		ExpiresAt: in.GetExpiresAt(),
	}

	ins, err := bson.Marshal(invalid)
	if err != nil {
		return err
	}

	collection := a.client.Collection("invalid_tokens")
	_, err = collection.InsertOne(ctx, ins)

	return err
}

func (a *Authorization) parseToken(in *authz.Token) (*jwt.Token, *claims.Claims, error) {
	c := &claims.Claims{}
	token, err := jwt.ParseWithClaims(in.Token, c, jwt.KnownKeyfunc(jwt.SigningMethodES256, a.privateKey))
	return token, c, err
}

func (a *Authorization) tokenInvalidated(ctx context.Context, tokenID *uuid.UUID) (bool, error) {
	var result *invalidToken
	filter := bson.M{"_id": tokenID}

	collection := a.client.Collection("invalid_tokens")
	if err := collection.FindOne(ctx, filter).Decode(result); err == mmongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
