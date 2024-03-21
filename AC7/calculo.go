package main

import (
	"fmt"
)

func calculo() {
	var code1, units1, price1 float64
	var code2, units2, price2 float64

	fmt.Scan(&code1, &units1, &price1)
	fmt.Scan(&code2, &units2, &price2)

	totalAmount := units1*price1 + units2*price2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totalAmount)
}
