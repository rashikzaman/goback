package domain

import "locally/goback/app/model"

type BusinessType struct {
	model.CommonModelFields
	Name string `json:"name" gorm:"not null"`
}
