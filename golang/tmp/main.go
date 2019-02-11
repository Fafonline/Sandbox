package main

import "fmt"

func main() {

	bareme := []float32{1000, 3000, 7000}

	var revenu float32 = 3000.0
	var tranche float32

	var tranchePrecedente float32 = 0
	for i := 0; i < 3; i++ {
		if i != 0 {
			tranchePrecedente = bareme[i-1]
		}
		if revenu <= bareme[i] {
			tranche = revenu - tranchePrecedente
			fmt.Printf("Tranche %d: %f\n", i, tranche)
			break
		} else {
			tranche = bareme[i] - tranchePrecedente
		}

		fmt.Printf("Tranche %d: %f\n", i, tranche)
	}

}
