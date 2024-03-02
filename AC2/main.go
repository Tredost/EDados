package main

import (
	"AC2/geometria"
	"fmt"
	"math"
	"math/rand"
)

func inverterString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type Ponto struct {
	X float64
	Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}

func main() {

	// EXERCICIO 1
	vetor := make([]int, 10)

	for i := 0; i < len(vetor); i++ {
		vetor[i] = rand.Intn(100)
	}

	for _, num := range vetor {
		fmt.Println(num)
	}

	// EXERCICIO 2
	fmt.Print("Digite uma string: ")
	var entrada string
	fmt.Scanln(&entrada)

	stringInvertida := inverterString(entrada)
	fmt.Println("String invertida:", stringInvertida)

	// EXERCICIO 3
	ponto1 := Ponto{3, 4}
	distancia := ponto1.DistanciaOrigem()
	fmt.Println("A distância do ponto (3, 4) até a origem é", distancia)

	ponto2 := Ponto{-2, 5}
	distancia = ponto2.DistanciaOrigem()
	fmt.Println("A distância do ponto (-2, 5) até a origem é", distancia)

	// EXERCICIO 4
	var base, altura float64

	fmt.Print("Digite a base do retângulo: ")
	fmt.Scanf("%f", &base)

	fmt.Print("Digite a altura do retângulo: ")
	fmt.Scanf("%f", &altura)
	fmt.Scanf("%f", &altura)

	retangulo := geometria.Retangulo{Base: base, Altura: altura}

	area := geometria.Area(retangulo)
	perimetro := geometria.Perimetro(retangulo)

	fmt.Println("A área do retângulo é", area)
	fmt.Println("O perímetro do retângulo é", perimetro)
}
