package repository

import (
	"context"
	"locally/goback/pkg/domain"
	"locally/goback/pkg/model"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	Conn *gorm.DB
}

func NewPostgresUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &PostgresUserRepository{Conn}
}

func (m *PostgresUserRepository) Fetch(ctx context.Context) model.Collection {
	var users []domain.User
	result := m.Conn.Find(&users)
	collection := model.Collection{
		Items:      users,
		PageNumber: 1,
		PageSize:   result.RowsAffected,
	}
	return collection
}

func (m *PostgresUserRepository) FetchById(ctx context.Context) {

}

func (m *PostgresUserRepository) FetchByEmail(ctx context.Context) {

}

func (m *PostgresUserRepository) Store(ctx context.Context) {

}

func (m *PostgresUserRepository) Update(ctx context.Context) {

}

func (m *PostgresUserRepository) Delete(ctx context.Context) {

}
