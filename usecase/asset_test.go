package usecase

import (
	"FlexcityTest/domain"
	"FlexcityTest/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestAsset_SelectAssetsForActivation(t *testing.T) {
	testCases := []struct {
		name         string
		request      domain.AssetsActivationRequest
		mockResponse []domain.Asset
		checks       func(*domain.AssetsActivationResponse, error)
	}{
		{
			name: "Sufficient assets available",
			request: domain.AssetsActivationRequest{
				Date:   time.Now(),
				Volume: 200,
			},
			mockResponse: []domain.Asset{
				{Code: "PUMP_1", ActivationCost: 100, Volume: 150},
				{Code: "PUMP_2", ActivationCost: 200, Volume: 100},
				{Code: "PUMP_3", ActivationCost: 9999, Volume: 200},
			},
			checks: func(response *domain.AssetsActivationResponse, err error) {
				require.NoError(t, err)
				assert.Equal(t, 2, len(response.Assets))
				assert.Equal(t, []string{"PUMP_1", "PUMP_2"}, []string{response.Assets[0].Code, response.Assets[1].Code})
				assert.Equal(t, 300, response.Price)
				assert.Equal(t, 250, response.Power)
			},
		},
		{
			name: "Sufficient assets available with multiples assets",
			request: domain.AssetsActivationRequest{
				Date:   time.Now(),
				Volume: 600,
			},
			mockResponse: []domain.Asset{
				{Code: "PUMP_1", ActivationCost: 100, Volume: 150},
				{Code: "PUMP_2", ActivationCost: 56, Volume: 75},
				{Code: "PUMP_3", ActivationCost: 48, Volume: 67},
				{Code: "PUMP_4", ActivationCost: 96, Volume: 141},
				{Code: "PUMP_5", ActivationCost: 3, Volume: 5},
				{Code: "PUMP_6", ActivationCost: 50, Volume: 150},
				{Code: "PUMP_7", ActivationCost: 200, Volume: 100},
			},
			checks: func(response *domain.AssetsActivationResponse, err error) {
				require.NoError(t, err)
				assert.Equal(t, 7, len(response.Assets))
				assert.Equal(t, 553, response.Price)
				assert.Equal(t, 688, response.Power)
			},
		},
		{
			name: "Insufficient assets available",
			request: domain.AssetsActivationRequest{
				Date:   time.Now(),
				Volume: 500,
			},
			mockResponse: []domain.Asset{
				{Code: "PUMP_1", ActivationCost: 100, Volume: 150},
				{Code: "PUMP_2", ActivationCost: 200, Volume: 100},
			},
			checks: func(response *domain.AssetsActivationResponse, err error) {
				require.Error(t, err)
				assert.Equal(t, domain.ErrNotEnoughAssets, err.(domain.ErrorResponse).Type)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.MockAssetRepository)
			usecase := assetUsecase{assetRepository: mockRepo}
			mockRepo.On("FindByAvailabilitySortedByCostRatio", tc.request.Date.Weekday()).Return(tc.mockResponse, nil)

			response, err := usecase.SelectAssetsForActivation(tc.request)
			tc.checks(response, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
