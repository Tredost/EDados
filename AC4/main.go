package main

import "fmt"

// EXERCÍCIO 1

func moverDisco(de, para string) {
	fmt.Printf("Mover o disco de %s para %s\n", de, para)
}

func hanoi(n int, de, para, aux string) {
	if n == 1 {
		moverDisco(de, para)
		return
	}
	hanoi(n-1, de, aux, para)
	moverDisco(de, para)
	hanoi(n-1, aux, para, de)
}

// EXERCÍCIO 2

func encontraPar(array []int, alvo int) (int, int) {
	esquerda := 0
	direita := len(array) - 1

	for esquerda < direita {
		soma := array[esquerda] + array[direita]
		if soma == alvo {
			return array[esquerda], array[direita]
		} else if soma < alvo {
			esquerda++
		} else {
			direita--
		}
	}

	return -1, -1
}

func main() {

	// EXERCÍCIO 1

	numeroDeDiscos := 3
	hanoi(numeroDeDiscos, "A", "C", "B")

	// EXERCÍCIO 2

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	alvo := 10
	num1, num2 := encontraPar(array, alvo)
	if num1 == -1 && num2 == -1 {
		fmt.Println("Nenhum par encontrado")
	} else {
		fmt.Printf("Par encontrado: %d e %d\n", num1, num2)
	}
}
