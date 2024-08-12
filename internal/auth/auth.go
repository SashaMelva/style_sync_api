package auth

import (
	"context"
)

type (
	AuthUser struct {
		Email    string
		Password string
	}
)

type Authentication interface {
	Authorization(ctx context.Context, user AuthUser) error
	Registration(ctx context.Context, new AuthUser) error
	UnAuthorization(ctx context.Context, user AuthUser) error
}
