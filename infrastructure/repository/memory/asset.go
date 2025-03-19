package memory

import (
	"FlexcityTest/domain"
	"FlexcityTest/infrastructure/repository"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"
)

func NewAssetRepositoryMemory() repository.AssetRepository {
	return assetRepositoryMemory{}
}

type assetRepositoryMemory struct {
}

// FindByAvailabilitySortedByCostRatio returns all assets available on the given weekday
// It prioritizes the assets with the lowest cost/volume ratio, and in case of a tie, it picks the one with the lowest activation cost
// In reality, with a database, it would be a query with a WHERE clause with ORDER BY on calculated field and then on activation cost, with pagination/limit
func (a assetRepositoryMemory) FindByAvailabilitySortedByCostRatio(weekday time.Weekday) ([]domain.Asset, error) {
	assets := getAllAssets()

	// Filter assets by availability
	var availableAssets []domain.Asset
	for _, asset := range assets {
		if slices.Contains(asset.Availability, weekday) {
			availableAssets = append(availableAssets, asset)
		}
	}

	// ASC Sort by cost/volume, cheaper first
	sort.Slice(availableAssets, func(first, second int) bool {
		firstRatio := getAssetRatioCostVolume(availableAssets[first])
		secondRatio := getAssetRatioCostVolume(availableAssets[second])
		if firstRatio == secondRatio {
			return availableAssets[first].ActivationCost < availableAssets[second].ActivationCost
		}

		return firstRatio < secondRatio
	})

	return availableAssets, nil
}

func getAssetRatioCostVolume(asset domain.Asset) float64 {
	return float64(asset.ActivationCost) / float64(asset.Volume)
}

func getAllAssets() []domain.Asset {
	allWeekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	allExceptWeekends := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday}
	onlySunday := []time.Weekday{time.Sunday}

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
		makeAsset(1, "Special", 750, onlySunday, 3500),
	}
}
