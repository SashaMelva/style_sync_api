package wardrobe

import (
	"context"

	"github.com/google/uuid"
)

type ApiConfig struct {
	WardrobeReader WardrobeReader
	WardrobeWriter WardrobeWriter
}

type Api struct {
	config ApiConfig
}

func NewApi(c ApiConfig) (*Api, error) {
	return &Api{
		config: ApiConfig{
			WardrobeReader: c.WardrobeReader,
			WardrobeWriter: c.WardrobeWriter},
	}, nil
}

func (api Api) GetAllWardrobe(ctx context.Context, userId uuid.UUID) (*Wardrobe, error) {
	return api.config.WardrobeReader.GetAllWardrobe(ctx, userId)
}
func (api Api) GetWardrobeByCategory(ctx context.Context, userId uuid.UUID, category string) (*Wardrobe, error) {
	return api.config.WardrobeReader.GetWardrobeByCategory(ctx, userId, category)
}

func (api Api) AddClotheToWardrobe(ctx context.Context, wardrobeId uuid.UUID, clothe Clothe) error {
	return api.config.WardrobeWriter.AddClotheToWardrobe(ctx, wardrobeId, clothe)
}
func (api Api) DeleteClotheToWardrobe(ctx context.Context, wardrobeId uuid.UUID, clotheId uuid.UUID) error {
	return api.config.WardrobeWriter.DeleteClotheToWardrobe(ctx, wardrobeId, clotheId)
}
func (api Api) AddWardrobe(ctx context.Context, wardrobe NewWardrobe) error {
	return api.config.WardrobeWriter.AddWardrobe(ctx, wardrobe)
}
func (api Api) UpdateWardrobe(ctx context.Context, wardrobe Wardrobe) error {
	return api.config.WardrobeWriter.UpdateWardrobe(ctx, wardrobe)
}
func (api Api) DeleteWardrobe(ctx context.Context, wardrobe uuid.UUID) error {
	return api.config.WardrobeWriter.DeleteWardrobe(ctx, wardrobe)
}
