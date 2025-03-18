package domain

type AlgorithmType string

const (
	RatioCostVolume AlgorithmType = "ratio_cost_volume"
	Backpack        AlgorithmType = "backpack"
)

var (
	AvailableAlgorithms = []AlgorithmType{RatioCostVolume, Backpack}
)
