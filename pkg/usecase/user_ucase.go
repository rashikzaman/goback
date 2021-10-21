package usecase

import (
	"context"
	"locally/goback/pkg/domain"
	"locally/goback/pkg/model"
)

type userUserCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(a domain.UserRepository) domain.UserUseCase {
	return &userUserCase{
		UserRepository: a,
	}
}

func (m *userUserCase) Fetch(ctx context.Context) model.Collection {
	collection := m.UserRepository.Fetch(ctx)
	return collection
}

func (m *userUserCase) FetchById(ctx context.Context) {

}

func (m *userUserCase) FetchByEmail(ctx context.Context) {

}

func (m *userUserCase) Store(ctx context.Context) {

}

func (m *userUserCase) Update(ctx context.Context) {

}

func (m *userUserCase) Delete(ctx context.Context) {

}
