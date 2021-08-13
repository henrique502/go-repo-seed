package team

import "time"

type Team struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New() Team {
	return Team{}
}
