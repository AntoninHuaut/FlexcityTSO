package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAsset_MarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		asset Asset
		want  string
	}{
		{
			name: "with decimal activation cost",
			asset: Asset{
				Code:           "code",
				Name:           "name",
				ActivationCost: 151,
				Availability:   []time.Weekday{time.Monday},
				Volume:         10,
			},
			want: `{"code":"code","name":"name","availability":[1],"volume":10,"activationCost":1.51}`,
		},
		{
			name: "with integer activation cost",
			asset: Asset{
				Code:           "code",
				Name:           "name",
				ActivationCost: 100,
				Availability:   []time.Weekday{},
				Volume:         10,
			},
			want: `{"code":"code","name":"name","availability":[],"volume":10,"activationCost":1}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.asset.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestAssetsActivationResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name               string
		activationResponse AssetsActivationResponse
		want               string
	}{
		{
			name: "with decimal price",
			activationResponse: AssetsActivationResponse{
				Assets: []Asset{},
				Price:  151,
				Power:  10,
			},
			want: `{"assets":[],"power":10,"price":1.51}`,
		},
		{
			name: "with integer price",
			activationResponse: AssetsActivationResponse{
				Assets: []Asset{},
				Price:  100,
				Power:  10,
			},
			want: `{"assets":[],"power":10,"price":1}`,
		},
		{
			name: "with assets",
			activationResponse: AssetsActivationResponse{
				Assets: []Asset{
					{
						Code:           "code1",
						Name:           "name1",
						ActivationCost: 100,
						Availability:   []time.Weekday{time.Sunday},
						Volume:         10,
					},
					{
						Code:           "code2",
						Name:           "name2",
						ActivationCost: 152,
						Availability:   []time.Weekday{},
						Volume:         10,
					},
				},
				Price: 9,
				Power: 10,
			},
			want: `{"assets":[{"code":"code1","name":"name1","availability":[0],"volume":10,"activationCost":1},{"code":"code2","name":"name2","availability":[],"volume":10,"activationCost":1.52}],"power":10,"price":0.09}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.activationResponse.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(got))
		})
	}
}
