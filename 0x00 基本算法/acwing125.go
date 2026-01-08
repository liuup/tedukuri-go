package main
import (
    "fmt"
    "sort"
)
type Cow struct {
    w int
    s int
}
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
func main() {
    var n int
    fmt.Scanf("%d", &n)
    cows := make([]Cow, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d %d", &cows[i].w, &cows[i].s)
    }
    sort.Slice(cows, func(i, j int) bool {
        return cows[i].w + cows[i].s < cows[j].w + cows[j].s
    })
    ans := -100000000
    sum := 0
    for i := 0; i < n; i++ {
        ans = max(ans, sum - cows[i].s)
        sum += cows[i].w
    }
    fmt.Println(ans)
}
