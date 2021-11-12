package usecase

import (
	"context"
	"locally/goback/app/domain"
	"locally/goback/app/model"
)

type businessUserCase struct {
	BusinessRepository domain.BusinessRepository
}

func NewBusinessUseCase(a domain.BusinessRepository) domain.BusinessUseCase {
	return &businessUserCase{
		BusinessRepository: a,
	}
}

func (m *businessUserCase) FetchBusinesses(ctx context.Context) (*model.Collection, error) {
	businesses, err := m.BusinessRepository.FetchBusinesses(ctx)
	if err != nil {
		return nil, err
	}
	collection := model.Collection{
		Items:      businesses,
		PageNumber: len(businesses),
		PageSize:   1,
	}
	return &collection, nil
}

func (m *businessUserCase) FetchBusinessById(ctx context.Context) {

}

func (m *businessUserCase) FetchByBusinessId(ctx context.Context) {

}

func (m *businessUserCase) StoreBusiness(ctx context.Context, data *domain.Business) (*domain.Business, error) {
	if data.Latitude != 0 && data.Longitude != 0 {
		data.Location = domain.Location{
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
		}
	}

	data, err := m.BusinessRepository.StoreBusiness(ctx, data)
	return data, err
}

func (m *businessUserCase) UpdateBusiness(ctx context.Context) {

}

func (m *businessUserCase) DeleteBusiness(ctx context.Context) {

}

func (m *businessUserCase) DeleteUser(ctx context.Context) {

}
