package pg

import (
	"context"

	"github.com/SashaMelva/style_sync_api/internal/profile"
	"github.com/google/uuid"
)

func (s *Storage) GetProfile(ctx context.Context, userId uuid.UUID) (*profile.Profile, error) {
	return nil, nil
}
func (s *Storage) UpdateProfile(ctx context.Context, profil profile.Profile) error {
	return nil
}
