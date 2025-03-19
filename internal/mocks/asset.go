package mocks

import (
	"FlexcityTest/domain"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockAssetRepository struct {
	mock.Mock
}

func (m *MockAssetRepository) FindByAvailabilitySortedByCostRatio(weekday time.Weekday) ([]domain.Asset, error) {
	args := m.Called(weekday)
	var assets []domain.Asset
	if args.Get(0) != nil {
		assets = args.Get(0).([]domain.Asset)
	}
	return assets, args.Error(1)
}

type MockAssetUsecase struct {
	mock.Mock
}

func (m *MockAssetUsecase) SelectAssetsForActivation(activationRequest domain.AssetsActivationRequest) (*domain.AssetsActivationResponse, error) {
	args := m.Called(activationRequest)
	var assets *domain.AssetsActivationResponse
	if args.Get(0) != nil {
		assets = args.Get(0).(*domain.AssetsActivationResponse)
	}
	return assets, args.Error(1)
}
