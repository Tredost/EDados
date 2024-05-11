package main

import "fmt"

func main() {
    var competidores, folhasCompradas, folhasPorCompetidor int
    fmt.Scan(&competidores, &folhasCompradas, &folhasPorCompetidor)

    if folhasCompradas >= competidores*folhasPorCompetidor {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}