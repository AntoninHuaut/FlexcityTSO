package repository

import (
	"FlexcityTest/domain"
	"time"
)

type AssetRepository interface {
	FindByAvailability(weekday time.Weekday) ([]domain.Asset, error)
}
