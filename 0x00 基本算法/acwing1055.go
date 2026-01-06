package main
import "fmt"
func max(a, b int) int {
    if (a > b) {
        return a
    }
    return b
}
func main() {
    var n, prev, ans int
    fmt.Scanf("%d\n%d", &n, &prev)
    for i := 1; i < n; i++ {
        var a int
        fmt.Scanf("%d", &a)
        ans += max(a - prev, 0)
        prev = a
    }
    fmt.Println(ans)
}
