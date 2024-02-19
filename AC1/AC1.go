package main

import "fmt"

func calculaMedia(numeros ...float64) float64 {
	var soma float64
	for _, numero := range numeros {
		soma += numero
	}
	return soma / float64(len(numeros))
}

func verificaParidade(numero int) bool {
	return numero%2 == 0
}

func minhaPotencia(base, expoente int) int {
	if expoente == 0 {
		return 1
	}
	resultado := 1
	for i := 0; i < expoente; i++ {
		resultado *= base
	}
	return resultado
}

func converteCelsiusParaFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func main() {
	media := calculaMedia(1.9, 5.2, 8.3, 2.4)
	fmt.Println("Média =", media)

	numero := 33
	par := verificaParidade(numero)
	if par {
		fmt.Println(numero, "é par")
	} else {
		fmt.Println(numero, "é ímpar")
	}

	potencia := minhaPotencia(3, 4)
	fmt.Println("3 elevado a 4 = ", potencia)

	temperaturaFahrenheit := converteCelsiusParaFahrenheit(30)
	fmt.Println("30 °C em Fahrenheit = ", temperaturaFahrenheit, "°F")
}
