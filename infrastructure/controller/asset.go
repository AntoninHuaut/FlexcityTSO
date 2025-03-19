package controller

import (
	"FlexcityTSO/domain"
	"FlexcityTSO/usecase"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

type AssetController interface {
	Activation(writer http.ResponseWriter, request *http.Request)
}

func NewAssetController(assetUsecase usecase.AssetUsecase) AssetController {
	return assetController{
		assetUsecase: assetUsecase,
	}
}

type assetController struct {
	assetUsecase usecase.AssetUsecase
}

func (a assetController) Activation(writer http.ResponseWriter, request *http.Request) {
	var activationRequest domain.AssetsActivationRequest

	if err := json.NewDecoder(request.Body).Decode(&activationRequest); err != nil {
		handleError(writer, request, domain.ErrorResponse{
			NativeError: err,
			Type:        domain.ErrInvalidPayload,
		})
		return
	}

	if formattedErr := validatePayload(activationRequest); formattedErr != nil {
		handleError(writer, request, formattedErr)
		return
	}

	activationResponse, formattedErr := a.assetUsecase.SelectAssetsForActivation(activationRequest)
	if formattedErr != nil {
		handleError(writer, request, formattedErr)
		return
	}

	render.JSON(writer, request, activationResponse)
}
