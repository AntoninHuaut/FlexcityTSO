package repository

import "FlexcityTest/domain"

type AssetRepository interface {
	FindAll() ([]domain.Asset, error)
}
