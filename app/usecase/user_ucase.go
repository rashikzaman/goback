package usecase

import (
	"context"
	"locally/goback/app/domain"
	"locally/goback/app/model"
	"time"
)

type userUserCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(a domain.UserRepository) domain.UserUseCase {
	return &userUserCase{
		UserRepository: a,
	}
}

func (m *userUserCase) FetchUsers(ctx context.Context) (*model.Collection, error) {
	result, err := m.UserRepository.FetchUsers(ctx)

	if err != nil {
		return nil, err
	}

	collection := model.Collection{
		Items:      result,
		PageNumber: 1,
		PageSize:   len(result),
	}
	return &collection, nil
}

func (m *userUserCase) FetchUserById(ctx context.Context) {

}

func (m *userUserCase) FetchUserByEmail(ctx context.Context) {

}

func (m *userUserCase) StoreUser(ctx context.Context, data *domain.User) (*domain.User, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	user, err := m.UserRepository.StoreUser(ctx, data)
	return user, err
}

func (m *userUserCase) UpdateUser(ctx context.Context) {

}

func (m *userUserCase) DeleteUser(ctx context.Context) {

}
