package dto

import "time"

type FunctionMinisterialResponseDTO struct {
	Id                    uint
	Name                  string
	Active                bool
	CreatedAt             time.Time
	UpdatedAt             *time.Time
	FunctionMinisterialId uint
}

type UserFunctionMinisterialDTO struct {
	Id                    uint `json:"id"`
	FunctionMinisterialId uint `json:"function_ministerial_id"`
	UserId                uint `json:"user_id"`
}
