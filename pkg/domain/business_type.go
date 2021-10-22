package domain

import "locally/goback/pkg/model"

type BusinessType struct {
	model.CommonModelFields
	Name string `json:"name" gorm:"not null"`
}
