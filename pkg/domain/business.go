package domain

import (
	"locally/goback/pkg/model"

	uuid "github.com/satori/go.uuid"
)

type Business struct {
	model.CommonModelFields
	BusinessId   uuid.UUID `gorm:"type:uuid;unique_index"`
	Name         string    `json:"name" gorm:"not null"`
	OwnerName    string    `json:"owner_name" gorm:"not null"`
	BusinessType BusinessType
	Street       string `json:"street_number"`
	Area         string `json:"area"`
	Thana        string `json:"thana"`
	District     string `json:"district"`
	Division     string `json:"division"`
	Lattitude    string `json:"lattitude"`
	Longitude    string `json:"longitude"`
}

type BusinessType struct {
	model.CommonModelFields
	Name string `json:"name" gorm:"not null"`
}
