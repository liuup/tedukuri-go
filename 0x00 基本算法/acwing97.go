package main

import "fmt"

func factorize(n int) map[int]int {
    factors := make(map[int]int)
    for n%2 == 0 && n > 0 {
        factors[2]++
        n /= 2
    }
    for i := 3; i*i <= n; i += 2 {
        for n%i == 0 {
            factors[i]++
            n /= i
        }
    }
    if n > 1 {
        factors[n]++
    }
    return factors
}

func fastpow(a, b, p int) int {
    ans := 1 % p
    for ; b != 0; b >>= 1 {
        if b & 1 == 1 {
            ans *= a
            ans %= p
        }
        a *= a
        a %= p
    }
    return ans
}

func sum(p, c int) int {
    if c == 0 {
        return 1
    }
    if c % 2 == 1 {
        return (1 + fastpow(p, (c+1)/2, 9901)) * sum(p, (c-1)/2) % 9901
    } else {
        return ((1 + fastpow(p, c/2, 9901)) * sum(p, c/2-1) % 9901 + fastpow(p, c, 9901)) % 9901
    }
}

func main() {
    var A, B int
    fmt.Scanf("%d %d", &A, &B)
    if A == 0 {
        fmt.Println(0)
        return
    }
    factors := factorize(A)
    ans := 1
    for k := range(factors) {
        factors[k] *= B
        ans *= sum(k, factors[k])
        ans %= 9901
    }
    fmt.Println(ans)
}
