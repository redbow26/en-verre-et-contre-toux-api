package dto

type CreateActivityDto struct {
	Name      string `json:"name" validate:"required"`
	MaxPlayer uint   `json:"max_player" validate:"required,min=1,max=6"`
}
