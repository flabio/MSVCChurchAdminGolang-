package entities

import "time"

type Church struct {
	Id        uint       `gorm:"primary_key:auto_increment" json:"id"`
	Name      string     `gorm:"type:varchar(250);not null" json:"name"`
	Phone     string     `gorm:"type:varchar(250);not null" json:"phone"`
	Email     string     `gorm:"type:varchar(250);not null" json:"email"`
	Address   string     `gorm:"type:varchar(100);not null" json:"address"`
	Active    bool       `gorm:"type:boolean" json:"active"`
	CreatedAt time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
