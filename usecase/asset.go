package usecase

import (
	"FlexcityTSO/domain"
	"FlexcityTSO/infrastructure/repository"
	"errors"
	"fmt"
)

type AssetUsecase interface {
	SelectAssetsForActivation(activationRequest domain.AssetsActivationRequest) (*domain.AssetsActivationResponse, error)
}

func NewAssetUsecase(assetRepository repository.AssetRepository) AssetUsecase {
	return assetUsecase{
		assetRepository: assetRepository,
	}
}

type assetUsecase struct {
	assetRepository repository.AssetRepository
}

// SelectAssetsForActivation selects the assets that will be activated based on the activation request
// returning the selected assets with total price and power, or an error if there are not enough assets to satisfy the demand
func (a assetUsecase) SelectAssetsForActivation(activationRequest domain.AssetsActivationRequest) (*domain.AssetsActivationResponse, error) {
	assets, err := a.assetRepository.FindByAvailabilitySortedByCostRatio(activationRequest.Date.Weekday())
	if err != nil {
		return nil, domain.ErrorResponse{
			NativeError: err,
			Type:        domain.ErrDatabase,
		}
	}

	var selectedAssets []domain.Asset
	totalVolume := 0
	totalCost := 0
	for _, asset := range assets {
		if totalVolume >= activationRequest.Volume {
			break
		}

		selectedAssets = append(selectedAssets, asset)
		totalCost += asset.ActivationCost
		totalVolume += asset.Volume
	}

	if totalVolume < activationRequest.Volume {
		return nil, domain.ErrorResponse{
			NativeError: errors.New(fmt.Sprintf("not enough available assets to satisfy the demand (%d)", activationRequest.Volume)),
			Type:        domain.ErrNotEnoughAssets,
		}
	}

	response := &domain.AssetsActivationResponse{
		Assets: selectedAssets,
		Price:  totalCost,
		Power:  totalVolume,
	}

	return response, nil
}
