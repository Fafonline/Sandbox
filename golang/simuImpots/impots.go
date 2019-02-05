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

	nombreTranchesImposable := c.dependencies.barême.Size()
	for indiceTrancheCourante := 0; indiceTrancheCourante < nombreTranchesImposable; indiceTrancheCourante++ {
		indiceTrancheSuivante := indiceTrancheCourante + 1
		if revenu > c.dependencies.barême.Get(indiceTrancheCourante) {
			tranches = ResizeSlice(tranches, indiceTrancheCourante+1)

			if indiceTrancheSuivante == nombreTranchesImposable {
				tranches[indiceTrancheCourante] = revenu - c.dependencies.barême.Get(indiceTrancheCourante)
				break
			}

			if revenu > c.dependencies.barême.Get(indiceTrancheSuivante) {
				tranches[indiceTrancheCourante] = c.dependencies.barême.Get(indiceTrancheSuivante) - c.dependencies.barême.Get(indiceTrancheCourante)
			} else {
				tranches[indiceTrancheCourante] = revenu - c.dependencies.barême.Get(indiceTrancheCourante)
			}
		} else {
			break
		}
	}

	return tranches
}
