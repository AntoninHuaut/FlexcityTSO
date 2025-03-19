package repository

import (
	"FlexcityTSO/domain"
	"time"
)

type AssetRepository interface {
	FindByAvailabilitySortedByCostRatio(weekday time.Weekday) ([]domain.Asset, error)
}
