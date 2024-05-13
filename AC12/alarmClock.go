package main

import (
    "fmt"
)

func main() {
    var H1, M1, H2, M2 int

    for {
        fmt.Scan(&H1, &M1, &H2, &M2)

        if H1 == 0 && M1 == 0 && H2 == 0 && M2 == 0 {
            break
        }

        currentMinutes := H1*60 + M1
        alarmMinutes := H2*60 + M2

        var sleepMinutes int

        if alarmMinutes > currentMinutes {
            sleepMinutes = alarmMinutes - currentMinutes
        } else {
            sleepMinutes = 24*60 - currentMinutes + alarmMinutes
        }

        fmt.Println(sleepMinutes)
    }
}
