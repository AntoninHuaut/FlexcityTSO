package boot

import (
	"FlexcityTSO/infrastructure/repository/memory"
	"FlexcityTSO/usecase"
)

var (
	assetUsecase usecase.AssetUsecase
)

func LoadServices() {
	assetRepository := memory.NewAssetRepositoryMemory()

	assetUsecase = usecase.NewAssetUsecase(assetRepository)
}
