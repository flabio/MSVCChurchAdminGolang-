package entities

type UserMinisterial struct {
	Id            uint        `gorm:"primary_key;auto_increment" `
	MinisterialId uint        `gorm:"null" `
	Ministerial   Ministerial `gorm:"foreignKey:MinisterialId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	UserId        uint        `gorm:"type:integer"`
}
