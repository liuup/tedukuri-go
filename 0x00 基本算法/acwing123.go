package main
import (
    "fmt"
    "sort"
    "math"
)

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    x, y := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d %d\n", &x[i], &y[i])
    }
    sort.Ints(x)
    sort.Ints(y)
    ans := 0
    for i := 0; i < n; i++ {
        x[i] -= i
    }
    sort.Ints(x)
    mdx := x[n/2]
    mdy := y[n/2]
    for i := 0; i < n; i++ {
        ans += int(math.Abs(float64(mdx - x[i])))
        ans += int(math.Abs(float64(mdy - y[i])))
    }
    fmt.Println(ans)
}
