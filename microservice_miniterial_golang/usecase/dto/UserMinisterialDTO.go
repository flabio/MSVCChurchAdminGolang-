package dto

type UserMinisterialDTO struct {
	Id            uint `json:"id"`
	MinisterialId uint `json:"ministerial_id"`
	UserId        uint `json:"user_id"`
}

type UsersMinisterialsDTO struct {
	Id            uint   `json:"id" `
	MinisterialId uint   `json:"ministerial_id"`
	UserId        uint   `json:"user_id"`
	Name          string `json:"name" validate:"required|min_len:3"  message:"required:The field {field} is required|min_len: min 3 caracters" label:"Name"`
	Active        bool   `json:"active" validate:"required" message:"required:{field} is required" label:"Active"`
}
