package repository

import (
	"context"
	"locally/goback/pkg/user/domain"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	Conn *gorm.DB
}

func NewPostgresUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &PostgresUserRepository{Conn}
}

func (m *PostgresUserRepository) Fetch(ctx context.Context) []domain.User {
	var users []domain.User
	m.Conn.Find(&users)
	return users
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
