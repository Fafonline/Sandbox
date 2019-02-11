package main

const (
	TRANCHE_01 float32 = 1000
	TRANCHE_02 float32 = 4000
	TRANCHE_03 float32 = 5000
)

type TrancheBuilderMock struct {
	TrancheBuilder
}

func (t TrancheBuilderMock) Build() []OutBoundTranche {
	return []OutBoundTranche{
		OutBoundTranche{
			threshold:   TRANCHE_01,
			coefficient: 0.05,
		},
		OutBoundTranche{
			threshold:   TRANCHE_02,
			coefficient: 0.1,
		},
		OutBoundTranche{
			threshold:   TRANCHE_03,
			coefficient: 0.30,
		},
	}
}
