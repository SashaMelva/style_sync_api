package auth

import "context"

type ApiConfig struct {
	Authentication Authentication
}

type Api struct {
	config Authentication
}

func NewApi(c ApiConfig) (*Api, error) {
	return &Api{
		config: c.Authentication,
	}, nil
}

func (api Api) Authorization(ctx context.Context, user AuthUser) error {
	return api.config.Authorization(ctx, user)
}
func (api Api) Registration(ctx context.Context, user AuthUser) error {
	return api.config.Registration(ctx, user)
}
func (api Api) UnAuthorization(ctx context.Context, user AuthUser) error {
	return api.config.UnAuthorization(ctx, user)
}
