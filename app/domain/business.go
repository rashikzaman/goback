package domain

import (
	"locally/goback/app/model"

	uuid "github.com/satori/go.uuid"

	"context"
)

type Business struct {
	model.CommonModelFields
	BusinessId     uuid.UUID `gorm:"type:uuid;unique_index"`
	Name           string    `json:"name" gorm:"not null"`
	OwnerName      string    `json:"owner_name" gorm:"not null"`
	BusinessTypeId uint      `gorm:"foreignKey:BusinessTypeId"`
	BusinessType   BusinessType
	Street         string `json:"street"`
	Area           string `json:"area"`
	Thana          string `json:"thana"`
	District       string `json:"district"`
	Division       string `json:"division"`
	Lattitude      string `json:"lattitude"`
	Longitude      string `json:"longitude"`
}

type BusinessRepository interface {
	FetchBusinesses(ctx context.Context) ([]Business, error)
	FetchBusinessById(ctx context.Context)
	FetchByBusinessId(ctx context.Context)
	StoreBusiness(ctx context.Context)
	UpdateBusiness(ctx context.Context)
	DeleteBusiness(ctx context.Context)
}

type BusinessUseCase interface {
	FetchBusinesses(ctx context.Context) (*model.Collection, error)
	FetchBusinessById(ctx context.Context)
	FetchByBusinessId(ctx context.Context)
	StoreBusiness(ctx context.Context)
	UpdateBusiness(ctx context.Context)
	DeleteBusiness(ctx context.Context)
}
