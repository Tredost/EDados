package main

import (
	"math"
	"fmt"
)

func mdc(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func ePrimo(num int) bool {
	if num <= 1 {
		return false
	}
	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func fatorial(n int64) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	resultado := int64(1)
	for i := int64(2); i <= n; i++ {
		resultado *= i
	}
	return resultado
}

func main() {
	var n, x, y int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		fmt.Println(mdc(x, y))
	}
	for {
		var x1, y1, x2, y2 int
		fmt.Scan(&x1, &y1, &x2, &y2)

		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		}

		if x1 == x2 && y1 == y2 {
			fmt.Println("0")
			continue
		}

		if x1 == x2 || y1 == y2 || math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) {
			fmt.Println("1")
			continue
		}

		fmt.Println("2")
	}
	var a, b int64
	for {
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			break
		}
		soma1 := fatorial(a)
		soma2 := fatorial(b)
		total := soma1 + soma2
		fmt.Println(total)
	}
	var q int
	fmt.Scan(&q)

	for i := 0; i < n; i++ {
		var num float64
		fmt.Scan(&num)

		dias := 0
		for num/2 > 1 {
			dias++
			num /= 2
		}

		fmt.Printf("%d dias\n", dias+1)
	}
	var qte, v int
	fmt.Scan(&qte)

	lista := make(map[int]int)

	for i := 0; i < qte; i++ {
		fmt.Scan(&v)
		lista[v]++
	}

	keys := make([]int, 0, len(lista))
	for k := range lista {
		keys = append(keys, k)
	}

	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	for _, k := range keys {
		fmt.Printf("%d aparece %d vez(es)\n", k, lista[k])
	}
	var m, l int
	fmt.Scan(&n)

	for i := 0; i < m; i++ {
		fmt.Scan(&l)
		if ePrimo(l) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
	}
}