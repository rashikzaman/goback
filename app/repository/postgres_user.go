package repository

import (
	"context"
	"locally/goback/app/domain"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	Conn *gorm.DB
}

func NewPostgresUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &PostgresUserRepository{Conn}
}

func (m *PostgresUserRepository) FetchUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	result := m.Conn.Find(&users)
	return users, result.Error
}

func (m *PostgresUserRepository) FetchUserById(ctx context.Context) {

}

func (m *PostgresUserRepository) FetchUserByEmail(ctx context.Context) {

}

func (m *PostgresUserRepository) StoreUser(ctx context.Context) {

}

func (m *PostgresUserRepository) UpdateUser(ctx context.Context) {

}

func (m *PostgresUserRepository) DeleteUser(ctx context.Context) {

}
