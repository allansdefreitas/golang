package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model        // it includes the control fields: ID, Created at, Updated at and Deleted at
	Name       string `json:"name" validate:"nonzero"`
	RG         string `json:"rg" validate:"len=7, regexp=^[0-9]*$"`
	CPF        string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
}

func ValidateStudent(student *Student) error {

	err := validator.Validate(student)
	if err != nil {
		return err
	}
	return nil

	// if err := validator.Validate(student); err != nil {
	// 	return err
	// }
	// return nil
}
