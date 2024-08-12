package profile

import (
	"context"

	"github.com/google/uuid"
)

type Profile struct {
	Id uuid.UUID
}

type ProfileReader interface {
	GetProfile(ctx context.Context, userId uuid.UUID) (*Profile, error)
}

type ProfileWriter interface {
	UpdateProfile(ctx context.Context, profil Profile) error
}
