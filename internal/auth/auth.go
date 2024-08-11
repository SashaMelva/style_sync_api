package auth

import (
	"context"

	"github.com/gofrs/uuid"
)

type (
	NewAuthUser struct {
		Email    string
		Password string
	}

	AuthUser struct {
		Id       uuid.UUID
		Email    string
		Password string
	}
)

type Authentication interface {
	Authorization(ctx context.Context, user AuthUser) error
	Registration(ctx context.Context, new NewAuthUser) error
	UnAuthorization(ctx context.Context, user AuthUser) error
}
