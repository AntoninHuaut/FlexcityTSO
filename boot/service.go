package boot

import (
	"FlexcityTest/infrastructure/repository/memory"
	"FlexcityTest/usecase"
)

var (
	assetUsecase usecase.AssetUsecase
)

func LoadServices() {
	assetRepository := memory.NewAssetRepositoryMemory()

	assetUsecase = usecase.NewAssetUsecase(assetRepository)
}
