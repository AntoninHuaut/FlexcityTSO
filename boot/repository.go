package boot

import "FlexcityTest/repository"

var (
	AssetRepository repository.AssetRepository
)

func LoadRepositories() {
	AssetRepository = repository.NewAssetRepositoryMemory()
}
