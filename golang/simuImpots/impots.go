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

	var tranchePrecedente float32 = 0
	for indexTrancheBarême := 0; indexTrancheBarême < nombreTranchesBarêmes; indexTrancheBarême++ {
		if indexTrancheBarême != 0 {
			tranchePrecedente = c.dependencies.barême.Get(indexTrancheBarême - 1)
		}
		tranches = ResizeSlice(tranches, indexTrancheBarême+1)
		if revenu <= c.dependencies.barême.Get(indexTrancheBarême) {
			tranches[indexTrancheBarême] = revenu - tranchePrecedente
			break
		} else {
			tranches[indexTrancheBarême] = c.dependencies.barême.Get(indexTrancheBarême) - tranchePrecedente
			if indexTrancheBarême == nombreTranchesBarêmes-1 {
				tranches = ResizeSlice(tranches, len(tranches)+1)
				tranches[indexTrancheBarême+1] = revenu - c.dependencies.barême.Get(indexTrancheBarême)
			}
		}
	}

	// for indiceTrancheCourante := 0; indiceTrancheCourante < nombreTranchesImposable; indiceTrancheCourante++ {
	// 	indiceTrancheSuivante := indiceTrancheCourante + 1
	// 	if revenu > c.dependencies.barême.Get(indiceTrancheCourante) {
	// 		tranches = ResizeSlice(tranches, indiceTrancheCourante+1)

	// 		if indiceTrancheSuivante == nombreTranchesImposable {
	// 			tranches[indiceTrancheCourante] = revenu - c.dependencies.barême.Get(indiceTrancheCourante)
	// 			break
	// 		}

	// 		if revenu > c.dependencies.barême.Get(indiceTrancheSuivante) {
	// 			tranches[indiceTrancheCourante] = c.dependencies.barême.Get(indiceTrancheSuivante) - c.dependencies.barême.Get(indiceTrancheCourante)
	// 		} else {
	// 			tranches[indiceTrancheCourante] = revenu - c.dependencies.barême.Get(indiceTrancheCourante)
	// 		}
	// 	} else {
	// 		break
	// 	}
	// }

	return tranches
}
