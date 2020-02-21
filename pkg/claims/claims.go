package claims

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/koverto/uuid"
)

type Claims struct {
	ID      *uuid.UUID `json:"jti,omitempty"`
	Subject *uuid.UUID `json:"sub,omitempty"`
	jwt.StandardClaims
}

func New(sub *uuid.UUID) *Claims {
	now := time.Now()
	exp := now.Add(30 * 24 * time.Hour)

	claims := &Claims{
		ID:      uuid.New(),
		Subject: sub,
	}

	claims.ExpiresAt = jwt.At(exp)
	claims.IssuedAt = jwt.At(now)
	claims.NotBefore = claims.IssuedAt

	return claims
}
