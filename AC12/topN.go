package main

import (
	"fmt"
)

func findCategory(position int) int {
	categories := []int{1, 3, 5, 10, 25, 50, 100}
	for _, category := range categories {
		if position <= category {
			return category
		}
	}
	return -1
}

func main() {
	var K int
	for {
		_, err := fmt.Scan(&K)
		if err != nil || K == 0 {
			break
		}
		category := findCategory(K)
		fmt.Printf("Top %d\n", category)
	}
}
