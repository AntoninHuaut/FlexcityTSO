package domain

import (
	"encoding/json"
	"time"
)

type Asset struct {
	Code           string         `json:"code"`
	Name           string         `json:"name"`
	ActivationCost int            `json:"-"`
	Availability   []time.Weekday `json:"availability"`
	Volume         int            `json:"volume"`
}

func (a Asset) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Asset
		ActivationCost float64 `json:"activationCost"`
	}{
		Asset:          a,
		ActivationCost: float64(a.ActivationCost) / 100.0,
	})
}
