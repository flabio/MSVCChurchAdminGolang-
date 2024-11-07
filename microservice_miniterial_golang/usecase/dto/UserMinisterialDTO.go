package dto

import "time"

type UserMinisterialDTO struct {
	Id            uint `json:"id"`
	MinisterialId uint `json:"ministerial_id"`
	UserId        uint `json:"user_id"`
}

type UsersMinisterialsDTO struct {
	Id            uint   `json:"id" `
	MinisterialId uint   `json:"ministerial_id"`
	UserId        uint   `json:"user_id"`
	Name          string `json:"name" `
	Active        bool   `json:"active"`
}

type MinisterialResponseDTO struct {
	Id            uint       `json:"id" `
	Name          string     `json:"name" `
	Active        bool       `json:"active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	MinisterialId uint       `json:"ministerial_id"`
}
