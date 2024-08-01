package dto

import (
	"microservice_user.com/usecase/client/church"
	"microservice_user.com/usecase/client/ministerial"
	"microservice_user.com/usecase/client/rol"
	"microservice_user.com/usecase/client/teams"
)

type UserDTO struct {
	FirstName          string `json:"first_name" validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"first_name"`
	LastName           string `json:"last_name" validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"last_name"`
	Identification     string `json:"identification"  validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"identification"`
	TypeIdentification string `json:"type_identification"`
	Birthdate          string `json:"birthdate" `
	PlaceOfBirth       string `json:"placeofbirth"`
	Address            string `json:"address"`
	Phone              string `json:"phone"`
	Cellphone          string `json:"cellphone"`
	Email              string `json:"email" validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"email"`
	Neighborhood       string `json:"neighborhood"`
	Locality           string `json:"locality"`
	Socioeconomic      uint   `json:"socioeconomic"`
	Rh                 string `json:"rh"`
	Profession         string `json:"profession"`
	Company            string `json:"company"`
	Position           string `json:"position"`
	CivilStatus        string `json:"civil_status"`
	LeadersName        string `json:"leaders_name"`
	ConversionDate     string `json:"conversion_date"`
	WhoInvitedHim      string `json:"who_invited_him"`
	ChurchAttended     bool   `json:"century_attended"`
	ChurchName         string `json:"church_name"`
	MeetingDate        string `json:"meetings_date"`
	BaptismDate        string `json:"baptism_date"`
	YourSpouseName     string `json:"your_spouse_name"`
	SpouseMemberChurch bool   `json:"spouse_member_church"`
	DateMarriage       string `json:"date_marriage"`
	ParentsName        string `json:"parents_name"`
	IsMemberParent     bool   `json:"is_member_parent"`
	ChildrenaName      string `json:"childrena_name"`
	Allergy            bool   `json:"allergy"`
	Authorize          bool   `json:"authorize"`
	Username           string `json:"username" validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"username"`
	Password           string `json:"password" validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"password"`
	ChurchId           uint   `json:"church_id" validate:"required" message:"required:The {field} is required " label:"church_id"`
	RolId              uint   `json:"rol_id" validate:"required" message:"required:The {field} is required " label:"rol_id"`
	TeamPescaId        uint   `json:"team_pesca_id" validate:"required" message:"required:The {field} is required " label:"team_pesca_id"`
}

type UserUpdateDTO struct {
	Id                 uint   `json:"id"`
	FirstName          string `json:"first_name" validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"first_name"`
	LastName           string `json:"last_name" validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"last_name"`
	Identification     string `json:"identification"  validate:"required|min_len:3" message:"required:The {field} is required|min_len: min 3 caracters" label:"identification"`
	TypeIdentification string `json:"type_identification" `
	Birthdate          string `json:"birthdate"`
	PlaceOfBirth       string `json:"placeofbirth"`
	Address            string `json:"address"`
	Phone              string `json:"phone"`
	Cellphone          string `json:"cellphone"`
	Email              string `json:"email"`
	Neighborhood       string `json:"neighborhood"`
	Locality           string `json:"locality"`
	Socioeconomic      uint   `json:"socioeconomic"`
	Rh                 string `json:"rh"`
	Profession         string `json:"profession"`
	Company            string `json:"company"`
	Position           string `json:"position"`
	CivilStatus        string `json:"civil_status"`
	LeadersName        string `json:"leaders_name"`
	ConversionDate     string `json:"conversion_date"`
	WhoInvitedHim      string `json:"who_invited_him"`
	ChurchAttended     bool   `json:"century_attended"`
	ChurchName         string `json:"church_name"`
	MeetingDate        string `json:"meetings_date"`
	BaptismDate        string `json:"baptism_date"`
	YourSpouseName     string `json:"your_spouse_name"`
	SpouseMemberChurch bool   `json:"spouse_member_church"`
	DateMarriage       string `json:"date_marriage"`
	ParentsName        string `json:"parents_name"`
	IsMemberParent     bool   `json:"is_member_parent"`
	ChildrenaName      string `json:"childrena_name"`
	Allergy            bool   `json:"allergy"`
	ChurchId           uint   `json:"church_id" validate:"required" message:"required:The {field} is required " label:"church_id"`
	RolId              uint   `json:"rol_id" validate:"required" message:"required:The {field} is required " label:"rol_id"`
	TeamPescaId        uint   `json:"team_pesca_id" validate:"required" message:"required:The {field} is required " label:"team_pesca_id"`
}

type UserReponse struct {
	Id              uint
	FirstName       string                        `json:"first_name"`
	LastName        string                        `json:"last_name"`
	Email           string                        `json:"email"`
	Phone           string                        `json:"phone"`
	Churchs         []church.Church               `json:"church"`
	Team            []teams.Team                  `json:"team"`
	Rol             []rol.Rol                     `json:"rol"`
	UserMinisterial []ministerial.UserMinisterial `json:"ministerial"`
}
