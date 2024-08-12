package pg

import (
	"context"

	"github.com/SashaMelva/style_sync_api/internal/wardrobe"
	"github.com/google/uuid"
)

func (s *Storage) GetAllWardrobe(ctx context.Context, userId uuid.UUID) (*wardrobe.Wardrobe, error) {
	return nil, nil
}
func (s *Storage) GetWardrobeByCategory(ctx context.Context, userId uuid.UUID, category string) (*wardrobe.Wardrobe, error) {
	return nil, nil
}

func (s *Storage) AddClotheToWardrobe(ctx context.Context, wardrobeId uuid.UUID, clothe wardrobe.Clothe) error {
	return nil
}
func (s *Storage) DeleteClotheToWardrobe(ctx context.Context, wardrobeId uuid.UUID, clotheId uuid.UUID) error {
	return nil
}
func (s *Storage) AddWardrobe(ctx context.Context, wardrobe wardrobe.NewWardrobe) error {
	return nil
}
func (s *Storage) UpdateWardrobe(ctx context.Context, wardrobe wardrobe.Wardrobe) error {
	return nil
}
func (s *Storage) DeleteWardrobe(ctx context.Context, wardrobe uuid.UUID) error {
	return nil
}
