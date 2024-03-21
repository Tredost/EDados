package main

import (
    "fmt"
)

func chimarrao() {
    var N int
    var L, Q float64
    fmt.Scan(&N, &L, &Q)

    participantes := make([]string, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&participantes[i])
    }

    index := int(L / Q) % N
    participanteRico := participantes[index]
    aguaRestante := L - float64(N-1)*Q

    fmt.Printf("%s %.1f\n", participanteRico, aguaRestante)
}
