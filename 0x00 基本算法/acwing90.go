package main

import (
    "fmt"
)

func main() {
    var a, b, mod int64
    fmt.Scanf("%d\n%d\n%d", &a, &b, &mod)

    fmt.Println(_64multiply(a, b, mod))
}

// 64位数乘法
func _64multiply(a, b, mod int64) (ans int64) {
    for ; b > 0; b >>= 1 {
        if b&1 == 1 {
            ans = (ans + a) % mod
        }
        a = a * 2 % mod
    }
    return
}
