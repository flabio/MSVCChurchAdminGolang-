package ministerial

import "time"

type Ministerial struct {
	Id            uint
	Name          string
	Active        bool
	MinisterialId uint
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

type UserMinisterial struct {
	Id            uint
	Ministerial   []Ministerial
	MinisterialId uint
	UserId        uint
}
