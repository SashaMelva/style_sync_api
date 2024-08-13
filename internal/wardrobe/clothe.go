package wardrobe

import "github.com/google/uuid"

type Clothe struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImgeUrl     string    `json:"image_url"`
	Params      string    `json:"params"`
}
