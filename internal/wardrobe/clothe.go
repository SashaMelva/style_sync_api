package wardrobe

import "github.com/google/uuid"

type Clothe struct {
	Id     uuid.UUID
	Title  string
	ImgUrl string
}
