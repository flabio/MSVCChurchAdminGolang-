package dto

type ChurchDTO struct {
	Id      uint   `json:"id"`
	Name    string `json:"name" `
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Active  bool   `json:"active"`
}
