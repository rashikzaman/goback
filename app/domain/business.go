package domain

import (
	"fmt"
	"locally/goback/app/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"context"
)

//https://gorm.io/docs/data_types.html
type Location struct {
	Longitude, Latitude float32
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", loc.Longitude, loc.Latitude)},
	}
}

func (b *Business) BeforeCreate(tx *gorm.DB) (err error) {
	b.BusinessId = uuid.NewV4()
	return
}

type Business struct {
	model.CommonModelFields
	BusinessId     uuid.UUID `gorm:"type:char(36);unique_index;not null"`
	Name           *string   `json:"name" gorm:"type:varchar(255);not null"`
	OwnerName      *string   `json:"owner_name" gorm:"type:varchar(255);not null"`
	BusinessTypeId *uint     `gorm:"foreignKey:BusinessTypeId;not null"`
	BusinessType   *BusinessType
	Street         string  `json:"street" gorm:"type:varchar(255)"`
	Area           string  `json:"area" gorm:"type:varchar(255)"`
	Thana          string  `json:"thana" gorm:"type:varchar(255)"`
	District       string  `json:"district" gorm:"type:varchar(255)"`
	Division       string  `json:"division" gorm:"type:varchar(255)"`
	Latitude       float32 `json:"latitude" gorm:"type:float;not null"`
	Longitude      float32 `json:"longitude" gorm:"type:float; not null"`
	Location       Location
}

type BusinessRepository interface {
	FetchBusinesses(ctx context.Context) ([]Business, error)
	FetchBusinessById(ctx context.Context)
	FetchByBusinessId(ctx context.Context)
	StoreBusiness(context.Context, *Business) (*Business, error)
	UpdateBusiness(ctx context.Context)
	DeleteBusiness(ctx context.Context)
}

type BusinessUseCase interface {
	FetchBusinesses(ctx context.Context) (*model.Collection, error)
	FetchBusinessById(ctx context.Context)
	FetchByBusinessId(ctx context.Context)
	StoreBusiness(context.Context, *Business) (*Business, error)
	UpdateBusiness(ctx context.Context)
	DeleteBusiness(ctx context.Context)
}
