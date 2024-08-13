package pg

import (
	"context"

	"github.com/SashaMelva/style_sync_api/internal/wardrobe"
	"github.com/google/uuid"
)

func (s *Storage) GetAllWardrobe(ctx context.Context, userId int) (*wardrobe.Wardrobe, error) {
	q := "SELECT id FROM wardrobe WHERE user_id = $1"
	var schema wardrobeSchema
	err := s.ConnectionDB.QueryRowContext(ctx, q, userId).Scan(&schema)
	if err != nil {
		return nil, err
	}

	res := wardrobe.Wardrobe{
		Id: schema.Id,
	}
	return &res, nil
}
func (s *Storage) GetWardrobeByCategory(ctx context.Context, userId int, category string) (*wardrobe.Wardrobe, error) {
	q := "SELECT FROM wardrobe WHERE user_id = $1 and category = $2"
	var schema wardrobeSchema
	err := s.ConnectionDB.QueryRowContext(ctx, q, userId).Scan(&schema)
	if err != nil {
		return nil, err
	}

	res := wardrobe.Wardrobe{
		Id: schema.Id,
	}
	return &res, nil
}

func (s *Storage) AddClotheToWardrobe(ctx context.Context, wardrobeId int, clothe wardrobe.Clothe) error {
	query := `insert into car_catalog(reg_num, mark, model, year, owner) values($1, $2, $3, $4, $5) RETURNING id`
	_, err := s.ConnectionDB.Exec(query, car.RegNum, car.Mark, car.Model, car.Year, car.Owner)

	if err != nil {
		return err
	}

	return nil
}
func (s *Storage) DeleteClotheToWardrobe(ctx context.Context, wardrobeId int, clotheId uuid.UUID) error {
	query := `delete from clothe_to_wardrobe where wardrobe_id = $1 and clothe_id=$2`
	_, err := s.ConnectionDB.Exec(query, wardrobeId, clotheId)

	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) AddWardrobe(ctx context.Context, wardrobe wardrobe.Wardrobe) error {
	query := `insert into wardrobe(reg_num, mark, model, year, owner) values($1, $2, $3, $4, $5) RETURNING id`
	_, err := s.ConnectionDB.Exec(query, car.RegNum, car.Mark, car.Model, car.Year, car.Owner)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateWardrobe(ctx context.Context, wardrobe wardrobe.Wardrobe) error {
	query := `update wardrobe set category=$1, title=$2, description=$3 where id = $4`
	_, err := s.ConnectionDB.Exec(query, wardrobe.Category, wardrobe.Title, wardrobe.Description, wardrobe.Id)

	if err != nil {
		return err
	}

	return nil
}
func (s *Storage) DeleteWardrobe(ctx context.Context, wardrobe int) error {
	query := `delete from wardrobe where id = $1`
	_, err := s.ConnectionDB.Exec(query, wardrobe)

	if err != nil {
		return err
	}
	return nil
}

type wardrobeSchema struct {
	Id          uuid.UUID `json:"id,omitempty"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Category    string    `json:"category"`
	// Clothes     []Clothe  `json:"clovers"`
}
