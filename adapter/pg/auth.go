package pg

import (
	"context"

	"github.com/SashaMelva/style_sync_api/internal/auth"
)

func (s *Storage) Authorization(ctx context.Context, user auth.AuthUser) error   { return nil }
func (s *Storage) Registration(ctx context.Context, new auth.NewAuthUser) error  { return nil }
func (s *Storage) UnAuthorization(ctx context.Context, user auth.AuthUser) error { return nil }
