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
	wardrobeReader WardrobeReader
	wardrobeWriter WardrobeWriter
}

func NewApi(c ApiConfig) (*Api, error) {
	return &Api{
		wardrobeReader: c.WardrobeReader,
		wardrobeWriter: c.WardrobeWriter,
	}, nil
}

func (api Api) GetAllWardrobe(ctx context.Context, userId uuid.UUID) (*Wardrobe, error) {
	return nil, nil
}
func (api Api) GetWardrobeByCategory(ctx context.Context, userId uuid.UUID, category string) (*Wardrobe, error) {
	return nil, nil
}

func (api Api) AddClotheToWardrobe(ctx context.Context, wardrobeId uuid.UUID, clothe Clothe) error {
	return nil
}
func (api Api) DeleteClotheToWardrobe(ctx context.Context, wardrobeId uuid.UUID, clotheId uuid.UUID) error {
	return nil
}
func (api Api) AddWardrobe(ctx context.Context, wardrobe NewWardrobe) error {
	return nil
}
func (api Api) UpdateWardrobe(ctx context.Context, wardrobe Wardrobe) error {
	return nil
}
func (api Api) DeleteWardrobe(ctx context.Context, wardrobe uuid.UUID) error {
	return nil
}
