package dto

type UpdateActivityDto struct {
	Name      string `json:"name" validate:""`
	MaxPlayer uint   `json:"max_player" validate:"omitempty,min=1,max=6"`
}
