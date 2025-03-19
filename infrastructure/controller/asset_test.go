package controller

import (
	"FlexcityTSO/domain"
	"FlexcityTSO/internal/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAsset_Activation(t *testing.T) {
	testCases := []struct {
		name         string
		request      func() ([]byte, error)
		prepareMocks func(*mocks.MockAssetUsecase)
		checks       func(*http.Response)
	}{
		{
			name: "nominal case",
			request: func() ([]byte, error) {
				return json.Marshal(domain.AssetsActivationRequest{
					Date:   time.Now().AddDate(0, 0, 1),
					Volume: 200,
				})
			},
			prepareMocks: func(mockUsecase *mocks.MockAssetUsecase) {
				mockUsecase.On("SelectAssetsForActivation", mock.Anything).Return(&domain.AssetsActivationResponse{
					Assets: []domain.Asset{{Code: "code", Name: "name", ActivationCost: 100, Availability: []time.Weekday{1}, Volume: 100}},
					Price:  100,
					Power:  100,
				}, nil)
			},
			checks: func(resp *http.Response) {
				require.Equal(t, http.StatusOK, resp.StatusCode)
			},
		},
		{
			name: "error case: nil request",
			request: func() ([]byte, error) {
				return nil, nil
			},
			prepareMocks: func(mockUsecase *mocks.MockAssetUsecase) {},
			checks: func(resp *http.Response) {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name: "error case: invalid payload",
			request: func() ([]byte, error) {
				return json.Marshal(domain.AssetsActivationRequest{
					Date:   time.Now().AddDate(0, 0, -1), // past date
					Volume: 200,
				})
			},
			prepareMocks: func(mockUsecase *mocks.MockAssetUsecase) {},
			checks: func(resp *http.Response) {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name: "error case: usecase generic error",
			request: func() ([]byte, error) {
				return json.Marshal(domain.AssetsActivationRequest{
					Date:   time.Now().AddDate(0, 0, 1),
					Volume: 200,
				})
			},
			prepareMocks: func(mockUsecase *mocks.MockAssetUsecase) {
				mockUsecase.On("SelectAssetsForActivation", mock.Anything).Return(nil, errors.New("generic error"))
			},
			checks: func(resp *http.Response) {
				require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
		{
			name: "error case: usecase no assets available error",
			request: func() ([]byte, error) {
				return json.Marshal(domain.AssetsActivationRequest{
					Date:   time.Now().AddDate(0, 0, 1),
					Volume: 200,
				})
			},
			prepareMocks: func(mockUsecase *mocks.MockAssetUsecase) {
				mockUsecase.On("SelectAssetsForActivation", mock.Anything).Return(nil, domain.ErrorResponse{Type: domain.ErrNotEnoughAssets})
			},
			checks: func(resp *http.Response) {
				require.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUsecase := new(mocks.MockAssetUsecase)
			controller := NewAssetController(mockUsecase)

			tc.prepareMocks(mockUsecase)

			jsonBody, err := tc.request()
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/v1/assets/activation", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			controller.Activation(w, req)
			resp := w.Result()

			tc.checks(resp)
			mockUsecase.AssertExpectations(t)
		})
	}
}
