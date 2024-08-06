package functionministerial

import "time"

type FunctionMinisterial struct {
	Id                    uint
	Name                  string
	Active                bool
	FunctionMinisterialId uint
	CreatedAt             time.Time
	UpdatedAt             *time.Time
}
