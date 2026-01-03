package main
import (
    "fmt"
    "sort"
    "math"
)
func main() {
    var n int
    fmt.Scanf("%d", &n)
    a := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &a[i])
    }
    sort.Ints(a)
    var middle int
    middle = a[n / 2]
    ans := 0
    for i := 0; i < n; i++ {
        ans += int(math.Abs(float64(a[i] - middle)))
    }
    fmt.Println(ans)
}
