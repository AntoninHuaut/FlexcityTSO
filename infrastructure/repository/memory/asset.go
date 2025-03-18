package memory

import (
	"FlexcityTest/domain"
	"FlexcityTest/infrastructure/repository"
	"fmt"
	"sort"
	"strings"
	"time"
)

func NewAssetRepositoryMemory() repository.AssetRepository {
	return assetRepositoryMemory{}
}

type assetRepositoryMemory struct {
}

// FindByAvailability returns all assets available on the given weekday (for database, it would be a query with a WHERE clause with ORDER BY)
func (a assetRepositoryMemory) FindByAvailability(weekday time.Weekday) ([]domain.Asset, error) {
	assets := getAllAssets()

	var availableAssets []domain.Asset
	for _, asset := range assets {
		for _, availability := range asset.Availability {
			if availability == weekday {
				availableAssets = append(availableAssets, asset)
				break
			}
		}
	}

	// ASC Sort by cost/volume
	sort.Slice(availableAssets, func(first, second int) bool {
		return availableAssets[first].ActivationCost/availableAssets[first].Volume < availableAssets[second].ActivationCost/availableAssets[second].Volume
	})

	return availableAssets, nil
}

func getAllAssets() []domain.Asset {
	allWeekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	allExceptWeekends := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday}

	makeAsset := func(number int, name string, activationCost int, availability []time.Weekday, volume int) domain.Asset {
		return domain.Asset{
			Code:           fmt.Sprintf("%s_%d", strings.ToUpper(name), number),
			Name:           fmt.Sprintf("%s %d", name, number),
			ActivationCost: activationCost,
			Availability:   availability,
			Volume:         volume,
		}
	}

	return []domain.Asset{
		makeAsset(1, "Pump", 1054, allWeekdays, 100),
		makeAsset(2, "Pump", 2850, allWeekdays, 300),
		makeAsset(3, "Pump", 3850, allWeekdays, 750),
		makeAsset(1, "Aeration", 1508, allWeekdays, 175),
		makeAsset(2, "Aeration", 2009, allExceptWeekends, 250),
		makeAsset(3, "Aeration", 5051, allExceptWeekends, 1500),
	}
}
