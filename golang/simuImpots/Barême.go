package main

type OutBoundTranche struct {
	threshold   float32
	coefficient float32
}

type Barême struct {
	outBoundTranche []OutBoundTranche
}

type Tranche struct {
	lowBound float32
	upBound  float32
	isLast   bool
	coeff    float32
}

func MakeBarême(outBoundTranche TrancheBuilder) (barême *Barême) {
	return &Barême{
		outBoundTranche: outBoundTranche.Build(),
	}
}

func (b *Barême) Get(index int) (tranche *Tranche) {

	var lowBound float32

	if index <= 0 {
		lowBound = 0
	} else {
		lowBound = b.outBoundTranche[index-1].threshold
	}
	upbound := b.outBoundTranche[index].threshold

	var isLast bool

	coeff := b.outBoundTranche[index].coefficient

	if index == b.Size()-1 {
		isLast = true
	} else {
		isLast = false
	}

	return &Tranche{
		lowBound: lowBound,
		upBound:  upbound,
		isLast:   isLast,
		coeff:    coeff,
	}
}

func (b *Barême) Size() int {
	return len(b.outBoundTranche)
}
