package entities

import "time"

type User struct {
	Id                 uint       `gorm:"primary key:auto_increment" json:"id"`
	FirstName          string     `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string     `gorm:"type:varchar(255);not null" json:"last_name"`
	Identification     string     `gorm:"type:varchar(255);not null" json:"identification"`
	TypeIdentification string     `gorm:"type:varchar(6);" json:"type_identification"`
	Birthdate          string     `gorm:"type:varchar(10)" json:"birthdate"`
	PlaceOfBirth       string     `gorm:"type:text" json:"place_of_birth"`
	Address            string     `gorm:"type:varchar(100)" json:"address"`
	Phone              string     `gorm:"type:varchar(20)" json:"phone"`
	Cellphone          string     `gorm:"type:varchar(20)" json:"cellphone"`
	Email              string     `gorm:"type:text" json:"email"`
	Neighborhood       string     `gorm:"type:text" json:"neighborhood"`
	Locality           string     `gorm:"type:text" json:"locality"`
	Socioeconomic      uint       `gorm:"type:integer" json:"socio_economic"`
	Rh                 string     `gorm:"type:varchar(6)" json:"rh"`
	Profession         string     `gorm:"type:text" json:"profession"`
	Company            string     `gorm:"type:text" json:"company"`
	Position           string     `gorm:"type:text" json:"position"`
	CivilStatus        string     `gorm:"type:varchar(50)" json:"civil_status"`
	LeadersName        string     `gorm:"type:text" json:"leader_name"`
	ConversionDate     string     `gorm:"type:varchar(10)" json:"conversion_date"`
	WhoInvitedHim      string     `gorm:"type:text" json:"who_invited_him"`
	ChurchAttended     bool       `gorm:"type:boolean" json:"church_attended"`
	ChurchName         string     `gorm:"type:varchar(250)" json:"church_name"`
	MeetingDate        string     `gorm:"type:varchar(10)" json:"meetings_date"`
	BaptismDate        string     `gorm:"type:varchar(10)" json:"baptism_date"`
	YourSpouseName     string     `gorm:"type:varchar(250)" json:"your_spouse_name"`
	SpouseMemberChurch bool       `gorm:"type:boolean" json:"spouse_member_church"`
	DateMarriage       string     `gorm:"type:varchar(10)" jon:"date_marriage"`
	ParentsName        string     `gorm:"type:varchar(250)" json:"parents_name"`
	IsMemberParent     bool       `gorm:"type:boolean" json:"is_member_parent"`
	ChildrenaName      string     `gorm:"type:varchar(250)" json:"childrena_name"`
	Allergy            bool       `gorm:"type:boolean" json:"allergy"`
	Authorize          bool       `gorm:"type:boolean" json:"authorize"`
	Username           string     `gorm:"type:string" json:"username"`
	Password           string     `gorm:"type:string" json:"password"`
	ChurchId           uint       `gorm:"type:integer" json:"church_id"`
	RolId              uint       `gorm:"type:integer" json:"rol_id"`
	TeamPescaId        uint       `gorm:"type:integer" json:"team_pesca_id"`
	CreatedAt          time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt          *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
