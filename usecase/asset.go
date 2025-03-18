package repository

import (
	"FlexcityTest/domain"
	"time"
)

type AssetRepository interface {
	FindAll() ([]domain.Asset, error)
}

func NewAssetRepository() AssetRepository {
	return assetRepository{}
}

type assetRepository struct {
}

func (a assetRepository) FindAll() ([]domain.Asset, error) {
	allWeekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	allExceptWeekends := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday}

	return []domain.Asset{
		{
			Code:           "PUMP_1",
			Name:           "Pump 1",
			ActivationCost: 1054,
			Availability:   allWeekdays,
			Volume:         100,
		},
		{
			Code:           "PUMP_2",
			Name:           "Pump 2",
			ActivationCost: 2850,
			Availability:   allWeekdays,
			Volume:         300,
		},
		{
			Code:           "AERATION_1",
			Name:           "Aeration 1",
			ActivationCost: 1500,
			Availability:   allWeekdays,
			Volume:         175,
		},
		{
			Code:           "AERATION_2",
			Name:           "Aeration 2",
			ActivationCost: 2000,
			Availability:   allExceptWeekends,
			Volume:         250,
		},
		{
			Code:           "AERATION_3",
			Name:           "Aeration 3",
			ActivationCost: 50,
			Availability:   allWeekdays,
			Volume:         3000,
		},
	}, nil
}
