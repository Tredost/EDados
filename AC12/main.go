package main

import (
	"fmt"
	"math"
)

func main() {
	var N, H, C, L int
	for {
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}
		fmt.Scan(&H, &C, &L)

		base_rampa := C * N
		altura_rampa := H * N
		largura_rampa := L

		comprimento_rampa := math.Sqrt(math.Pow(float64(altura_rampa), 2) + math.Pow(float64(base_rampa), 2))

		area_rampa := float64(largura_rampa) * comprimento_rampa

		fmt.Printf("%.4f\n", area_rampa/10000)
	}
}
