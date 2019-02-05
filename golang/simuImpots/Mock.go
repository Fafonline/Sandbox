package main

const (
	TRANCHE_01 float32 = 1000
	TRANCHE_02 float32 = 4000
	TRANCHE_03 float32 = 5000
)

type TrancheBuilderMock struct {
	TrancheBuilder
}

func (t TrancheBuilderMock) Build() []float32 {
	return []float32{TRANCHE_01, TRANCHE_02, TRANCHE_03}
}
