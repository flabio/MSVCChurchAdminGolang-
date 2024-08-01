package ministerial

type Ministerial struct {
	Id     uint
	Name   string
	Active bool
}

type UserMinisterial struct {
	Id            uint
	MinisterialId uint
	UserId        uint
	Ministerial   []Ministerial
}
