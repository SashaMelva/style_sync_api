package auth

import "context"

type ApiConfig struct {
	Authentication Authentication
}

type Api struct {
	authentication Authentication
}

func NewApi(c ApiConfig) (*Api, error) {
	return &Api{
		authentication: c.Authentication,
	}, nil
}

func (api Api) Authorization(ctx context.Context, user AuthUser) error   { return nil }
func (api Api) Registration(ctx context.Context, new NewAuthUser) error  { return nil }
func (api Api) UnAuthorization(ctx context.Context, user AuthUser) error { return nil }
