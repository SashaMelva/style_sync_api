package wardrobe

import (
	"context"

	"github.com/google/uuid"
)

type Wardrobe struct {
	Id          uuid.UUID `json:"id,omitempty"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Category    string    `json:"category"`
	Clothes     []Clothe  `json:"clovers"`
}

type WardrobeReader interface {
	GetAllWardrobe(ctx context.Context, userId uuid.UUID) (*Wardrobe, error)
	GetWardrobeByCategory(ctx context.Context, userId uuid.UUID, category string) (*Wardrobe, error)
}

type WardrobeWriter interface {
	AddClotheToWardrobe(ctx context.Context, wardrobeId uuid.UUID, clothe Clothe) error
	DeleteClotheToWardrobe(ctx context.Context, wardrobeId uuid.UUID, clotheId uuid.UUID) error

	AddWardrobe(ctx context.Context, wardrobe Wardrobe) error
	UpdateWardrobe(ctx context.Context, wardrobe Wardrobe) error
	DeleteWardrobe(ctx context.Context, wardrobe uuid.UUID) error
}
