package model

type User struct {
	*Base
	*UserInput
}

type UserInput struct {
	FirstName string `json:"firstName" xml:"firstName" gorm:"column:first_name;size:255;not null" validate:"required"`
	LastName  string `json:"lastName" xml:"lastName" gorm:"column:last_name;size:255;not null" validate:"required"`
	Email     string `json:"email" xml:"email" gorm:"column:email;size:255;not null" validate:"required,email"`
}
