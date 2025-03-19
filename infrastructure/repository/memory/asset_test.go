package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAsset_FindByAvailability(t *testing.T) {
	repo := assetRepositoryMemory{}

	testCases := []struct {
		weekday  time.Weekday
		expected []string
	}{
		{time.Monday, []string{"AERATION_3", "PUMP_3", "AERATION_2", "AERATION_1", "PUMP_2", "PUMP_1"}},
		{time.Saturday, []string{"PUMP_3", "AERATION_1", "PUMP_2", "PUMP_1"}},
		{time.Sunday, []string{"SPECIAL_1", "PUMP_3", "AERATION_1", "PUMP_2", "PUMP_1"}},
	}

	for _, tc := range testCases {
		assets, err := repo.FindByAvailabilitySortedByCostRatio(tc.weekday)
		assert.NoError(t, err)
		var assetCodes []string
		for _, asset := range assets {
			assetCodes = append(assetCodes, asset.Code)
		}
		assert.Equal(t, tc.expected, assetCodes)
	}
}
