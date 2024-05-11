package main

import "fmt"

func main() {
    var distance int
    fmt.Scanln(&distance)

    timeTaken := float64(distance) / 30 * 60

    fmt.Printf("%.0f minutos\n", timeTaken)
}
