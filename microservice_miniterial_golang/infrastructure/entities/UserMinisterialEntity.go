package entities

type UserMinisterial struct {
	Id            uint        `gorm:"primary_key;auto_increment" json:"id"`
	MinisterialId uint        `gorm:"null" json:"ministerial_id"`
	Ministerial   Ministerial `gorm:"foreignKey:MinisterialId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"ministerials"`
	UserId        uint        `gorm:"type:integer" json:"user_id"`
}
