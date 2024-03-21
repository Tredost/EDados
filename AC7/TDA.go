package main

import (
    "fmt"
)

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func simplify(numerator, denominator int) (int, int) {
    common := gcd(numerator, denominator)
    return numerator / common, denominator / common
}

func TDA() {
    var N int
    fmt.Scan(&N)

    for i := 0; i < N; i++ {
        var N1, D1, N2, D2 int
        var op rune
        fmt.Scan(&N1, &op, &D1, &N2, &D2)

        var resultN, resultD int

        switch op {
        case '+':
            resultN = N1*D2 + N2*D1
            resultD = D1 * D2
        case '-':
            resultN = N1*D2 - N2*D1
            resultD = D1 * D2
        case '*':
            resultN = N1 * N2
            resultD = D1 * D2
        case '/':
            resultN = N1 * D2
            resultD = N2 * D1
        }

        simplifiedN, simplifiedD := simplify(resultN, resultD)

        fmt.Printf("%d/%d = %d/%d\n", resultN, resultD, simplifiedN, simplifiedD)
    }
}
