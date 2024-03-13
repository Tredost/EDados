package main

import (
	"fmt"
	"math"
	"sort"
)

func exercicio2() {

	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	sides := []float64{A, B, C}
	sort.Float64s(sides)
	A, B, C = sides[2], sides[1], sides[0]

	if A >= B+C {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if math.Pow(A, 2) == math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO RETANGULO")
		}
		if math.Pow(A, 2) > math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if math.Pow(A, 2) < math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO ACUTANGULO")
		}
		if A == B && B == C {
			fmt.Println("TRIANGULO EQUILATERO")
		}
		if (A == B && A != C) || (A == C && A != B) || (B == C && B != A) {
			fmt.Println("TRIANGULO ISOSCELES")
		}
	}
}
