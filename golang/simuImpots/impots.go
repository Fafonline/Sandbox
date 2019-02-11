package main

type Dependencies struct {
	barême *Barême
}

type CalculateurTranche struct {
	dependencies *Dependencies
}

func NewCalculateurTranche(dependencies *Dependencies) *CalculateurTranche {
	return &CalculateurTranche{
		dependencies: dependencies,
	}
}

func (c *CalculateurTranche) Execute(revenu float32) (tranches []float32) {

	tranches = nil
	nombreTranchesBarêmes := c.dependencies.barême.Size()

	for indexTrancheBarême := 0; indexTrancheBarême < nombreTranchesBarêmes; indexTrancheBarême++ {

		tranche := c.dependencies.barême.Get(indexTrancheBarême)
		lowBound := tranche.lowBound
		upbound := tranche.upBound
		isLast := tranche.isLast

		tranches = ResizeSlice(tranches, indexTrancheBarême+1)
		if revenu <= upbound {
			tranches[indexTrancheBarême] = revenu - lowBound
			break
		} else {
			tranches[indexTrancheBarême] = upbound - lowBound
			if isLast {
				tranches = ResizeSlice(tranches, len(tranches)+1)
				tranches[indexTrancheBarême+1] = revenu - upbound
			}
		}
	}

	return tranches
}
