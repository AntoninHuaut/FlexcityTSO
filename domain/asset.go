package domain

import (
	"encoding/json"
	"time"
)

type Asset struct {
	Code           string         `json:"code"`
	Name           string         `json:"name"`
	ActivationCost int            `json:"-"` // ActivationCost in cents
	Availability   []time.Weekday `json:"availability"`
	Volume         int            `json:"volume"`
}

// MarshalJSON is a custom JSON marshaller for Asset to convert ActivationCost to a float64
func (a Asset) MarshalJSON() ([]byte, error) {
	type Alias Asset
	return json.Marshal(&struct {
		*Alias
		ActivationCost float64 `json:"activationCost"`
	}{
		Alias:          (*Alias)(&a),
		ActivationCost: float64(a.ActivationCost) / 100.0,
	})
}

type AssetsActivationRequest struct {
	Date   time.Time `validate:"required,notBeforeNow" json:"date"`
	Volume int       `validate:"required,min=1" json:"volume"`
}

type AssetsActivationResponse struct {
	Assets []Asset `json:"assets"`
	Price  int     `json:"price"` // Price in cents
	Power  int     `json:"power"`
}

// MarshalJSON is a custom JSON marshaller for AssetsActivationResponse to convert Price to a float64
func (a AssetsActivationResponse) MarshalJSON() ([]byte, error) {
	type Alias AssetsActivationResponse
	return json.Marshal(&struct {
		*Alias
		Price float64 `json:"price"`
	}{
		Alias: (*Alias)(&a),
		Price: float64(a.Price) / 100.0,
	})
}
