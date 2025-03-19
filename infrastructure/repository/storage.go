package repository

import (
	"FlexcityTest/domain"
	"time"
)

type AssetRepository interface {
	FindByAvailabilitySortedByCostRatio(weekday time.Weekday) ([]domain.Asset, error)
}
