package ministerial

import "time"

type Ministerial struct {
	Active    bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	Id        uint
	Name      string
}

type UserMinisterial struct {
	Id            uint
	Ministerial   []Ministerial
	MinisterialId uint
	UserId        uint
}
