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
	profileReader ProfileReader
	profileWriter ProfileWriter
}

func NewApi(c ApiConfig) (*Api, error) {
	return &Api{
		profileReader: c.ProfileReader,
		profileWriter: c.ProfileWriter,
	}, nil
}

func (api Api) GetProfile(ctx context.Context, userId uuid.UUID) (*Profile, error) {
	return nil, nil
}

func (api Api) UpdateProfile(ctx context.Context, profil Profile) error {
	return nil
}
