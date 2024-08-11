package profile

import (
	"context"

	"github.com/google/uuid"
)

type Profile struct {
}

type ProfileReader interface {
	GetProfile(ctx context.Context, userId uuid.UUID) (*Profile, error)
}

type ProfileWriter interface {
	UpdateProfile(ctx context.Context, profil Profile) error
}
