package main

import "fmt"

func main() {
    var n int
    var jumlah int
    var t1, t2, t3 int

    fmt.Scanln(&n)
    for i := 0; i < n; i++ {
        fmt.Scanln(&t1, &t2, &t3)

        if int(t1)+int(t2)+int(t3) >= 2 {
            jumlah += 1
        }
    }
    fmt.Println(jumlah)
}