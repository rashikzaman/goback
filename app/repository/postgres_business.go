package repository

import (
	"context"
	"locally/goback/app/domain"

	"gorm.io/gorm"
)

type PostgresBusinessRepository struct {
	Conn *gorm.DB
}

func NewPostgresBusinessRepository(Conn *gorm.DB) domain.BusinessRepository {
	return &PostgresBusinessRepository{Conn}
}

func (m *PostgresBusinessRepository) FetchBusinesses(ctx context.Context) ([]domain.Business, error) {
	var businesses []domain.Business
	result := m.Conn.Find(&businesses)
	return businesses, result.Error
}

func (m *PostgresBusinessRepository) FetchByBusinessId(ctx context.Context) {

}

func (m *PostgresBusinessRepository) FetchBusinessById(ctx context.Context) {

}

func (m *PostgresBusinessRepository) FetchBusinessByEmail(ctx context.Context) {

}

func (m *PostgresBusinessRepository) StoreBusiness(ctx context.Context) {

}

func (m *PostgresBusinessRepository) UpdateBusiness(ctx context.Context) {

}

func (m *PostgresBusinessRepository) DeleteBusiness(ctx context.Context) {

}
