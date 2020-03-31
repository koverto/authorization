// Package claims handles generating and parsing JWT claims.
package claims

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/koverto/uuid"
)

// ContextKeyJTI is the key used for storing JWT IDs in a context.
type ContextKeyJTI struct{}

// ContextKeySUB is the key used for storing JWT subjects in a context.
type ContextKeySUB struct{}

// ContextKeyEXP is the key used for storing JWT expiration timestamps in a context.
type ContextKeyEXP struct{}

// Claims defines the structure of JWT claims.
type Claims struct {
	ID      *uuid.UUID `json:"jti,omitempty"`
	Subject *uuid.UUID `json:"sub,omitempty"`
	jwt.StandardClaims
}

// New generates a new set of JWT claims for the given subject.
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
