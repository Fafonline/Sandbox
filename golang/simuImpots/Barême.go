package main

type Barême struct {
	tranches []float32
}

type Tranche struct {
	lowBound float32
	upBound  float32
	isLast   bool
}

func (b *Barême) Error() string {
	return "Erreur Barême"
}

type BarêmeError struct {
	error
	errorString string
}

func MakeBarêmeError(errorString string) *BarêmeError {
	return &BarêmeError{
		errorString: errorString,
	}
}
func (br *BarêmeError) Error() string {
	return br.errorString
}

func MakeBarême(tranches TrancheBuilder) (barême *Barême) {
	return &Barême{
		tranches: tranches.Build(),
	}
}

func (b *Barême) Get(index int) (tranche *Tranche) {

	var lowBound float32

	if index <= 0 {
		lowBound = 0
	} else {
		lowBound = b.tranches[index-1]
	}
	upbound := b.tranches[index]

	var isLast bool

	if index == b.Size()-1 {
		isLast = true
	} else {
		isLast = false
	}

	return &Tranche{
		lowBound: lowBound,
		upBound:  upbound,
		isLast:   isLast,
	}
}

func (b *Barême) Size() int {
	return len(b.tranches)
}

func (b *Barême) TrancheComplete(indiceTranche int) (tranche float32, err error) {
	indiceTrancheSuivante := indiceTranche + 1
	if (indiceTrancheSuivante) == b.Size() {
		return 0, MakeBarêmeError("Tranche Maximale")
	} else {
		return b.tranches[indiceTrancheSuivante] - b.tranches[indiceTranche], nil
	}
}
