package main

import (
    "fmt"
)

func main() {
    var trips int
    fmt.Scanln(&trips)

    for i := 0; i < trips; i++ {
        var products, shoppingList int
        fmt.Scanln(&products)

        // Create a map to store prices of products
        prices := make(map[string]float64)

        // Read products and their prices
        for j := 0; j < products; j++ {
            var productName string
            var price float64
            fmt.Scan(&productName, &price)
            prices[productName] = price
        }

        fmt.Scanln(&shoppingList)

        // Calculate the total cost for the shopping list
        var totalCost float64
        for j := 0; j < shoppingList; j++ {
            var productName string
            var quantity int
            fmt.Scan(&productName, &quantity)
            totalCost += prices[productName] * float64(quantity)
        }

        // Print the total cost for this test case
        fmt.Printf("R$ %.2f\n", totalCost)
    }
}
