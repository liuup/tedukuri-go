package main
import (
    "fmt"
    "sort"
    "math"
)
type Range struct {
    min float64
    max float64
}
func min(a, b float64) float64 {
    if (a < b) {
        return a
    }
    return b
}
func main() {
    var n, d int
    fmt.Scanf("%d %d", &n, &d)
    ranges := make([]*Range, 0, n)
    for i := 0; i < n; i++ {
        var x, y int
        fmt.Scanf("%d %d", &x, &y)
        if d*d-y*y < 0{
            fmt.Println(-1)
            return
        }
        ranges = append(ranges, &Range{
            min: float64(x)-math.Sqrt(float64(d*d-y*y)),
            max: float64(x)+math.Sqrt(float64(d*d-y*y)),
        })
    }
    sort.Slice(ranges, func(i, j int) bool {
        return ranges[i].min < ranges[j].min
    })
    last := -2e9
    ans := 0
    for _, rng := range ranges {
        if rng.min > rng.max {
            ans = -1
            break
        }
        if rng.min <= last {
            last = min(last, rng.max)
        } else {
            ans++
            last = rng.max
        }
    }
    fmt.Println(ans)
}
