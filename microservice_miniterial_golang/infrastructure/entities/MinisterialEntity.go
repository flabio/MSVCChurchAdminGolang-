package entities

import "time"

type Ministerial struct {
	Id              uint               `gorm:"primary_key:auto_increment" `
	Name            string             `gorm:"type:varchar(100);not null" json:"name"`
	Active          bool               `gorm:"type:boolean" json:"active"`
	CreatedAt       time.Time          `gorm:"<-created_at" json:"created_at"`
	UpdatedAt       *time.Time         `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	UserMinisterial *[]UserMinisterial `json:"user_ministerial,omitempty"`
}
