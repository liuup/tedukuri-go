package main
import "fmt"
var n, f int
var a [100005]float64
var sum [100005]float64
func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}
func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
func check(val float64) bool {
    min_val := 10000000000.0
    ans := -10000000000.0
    for i := 1; i <= n; i++ {
        sum[i] = sum[i-1] + a[i] - val
    }
    for i := f; i <= n; i++ {
        min_val = min(min_val, sum[i-f])
        ans = max(ans, sum[i] - min_val)
    }
    return ans >= 0
}
func main() {
    fmt.Scanf("%d %d", &n, &f)
    for i := 1; i <= n; i++ {
        fmt.Scanf("%f", &a[i])
    }
    l := -1e6
    r := 1e6
    for l + 1e-6 < r {
        mid := (l + r) / 2
        if check(mid) {
            l = mid
        } else {
            r = mid
        }
    }
    fmt.Println(int(1000 * r))
}
