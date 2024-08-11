package schema

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Location string    `json:"location"`
}

type NewUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
