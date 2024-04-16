package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model        // it includes the control fields: ID, Created at, Updated at and Deleted at
	Name       string `json:"name"`
	CPF        string `json:"cpf"`
	RG         string `json:"rg"`
}
