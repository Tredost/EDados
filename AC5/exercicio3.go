package main

import (
	"fmt"
)

func exercicio3() {
	var T int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		var PA, PB int
		var G1, G2 float64
		fmt.Scanln(&PA, &PB, &G1, &G2)

		anos := 0
		for PA <= PB {
			PA = int(float64(PA) * (1 + G1/100))
			PB = int(float64(PB) * (1 + G2/100))
			anos++

			if anos > 100 {
				fmt.Println("Mais de 1 seculo")
				break
			}
		}

		if anos <= 100 {
			fmt.Printf("%d anos.\n", anos)
		}
	}
}
