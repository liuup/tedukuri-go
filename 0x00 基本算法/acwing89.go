package main

import "fmt"

func fastpow(a, b, p int64) int64 {
    var ans int64
	ans = 1 % p
	for ; b > 0; b >>= 1 {
		if b&1 != 0 {
			ans *= a
			ans %= p
		}
		a *= a
		a %= p
	}
	return ans
}

func main() {
	var a, b, p int64
	fmt.Scanf("%d %d %d", &a, &b, &p)
	fmt.Println(fastpow(a, b, p))
}
