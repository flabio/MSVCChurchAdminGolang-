package entities

type UserFunctionMinisterial struct {
	Id                    uint                `gorm:"primary_key:auto_increment"`
	FunctionMinisterialId uint                `gorm:"null"`
	FunctionMinisterial   FunctionMinisterial `gorm:"foreignKey:FunctionMinisterialId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	UserId                uint                `gorm:"type:integer"`
}
