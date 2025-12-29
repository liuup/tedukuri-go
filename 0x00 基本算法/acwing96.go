package main

import "fmt"

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    var d, f [13]int
    d[0] = 0
    f[0] = 0
    for i := 1; i <= 12; i++ {
        d[i] = 2 * d[i-1] + 1
        f[i] = d[i]
        for j := 1; j < i; j++ {
            f[i] = min(f[i], 2 * f[j] + d[i-j])
        }
        fmt.Println(f[i])
    }
}
