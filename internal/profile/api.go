package profile

import (
	"context"

	"github.com/google/uuid"
)

type ApiConfig struct {
	ProfileReader ProfileReader
	ProfileWriter ProfileWriter
}

type Api struct {
	config ApiConfig
}

func NewApi(c ApiConfig) (*Api, error) {
	return &Api{
		config: ApiConfig{
			ProfileReader: c.ProfileReader,
			ProfileWriter: c.ProfileWriter},
	}, nil
}

func (api Api) GetProfile(ctx context.Context, userId uuid.UUID) (*Profile, error) {
	return api.config.ProfileReader.GetProfile(ctx, userId)
}

func (api Api) UpdateProfile(ctx context.Context, profil Profile) error {
	return api.config.ProfileWriter.UpdateProfile(ctx, profil)
}
