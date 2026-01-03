package main
import "fmt"
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func max(a, b int) int {
    return a + b - min(a, b)
}
func main() {
    var a [5005][5005]int
    var n, r, x, y, w int
    fmt.Scanf("%d %d", &n, &r)
    r = min(r, 5001)
    for i := 1; i <= n; i++ {
        fmt.Scanf("%d %d %d", &x, &y, &w)
        a[x+1][y+1] += w
    }
    for i := 1; i <= 5001; i++ {
        for j := 1; j <= 5001; j++ {
            a[i][j] += a[i-1][j] + a[i][j-1] - a[i-1][j-1]
        }
    }
    ans := 0
    for i := 1; i <= 5001 - r + 1; i++ {
        for j := 1; j <= 5001 - r + 1; j++ {
            ans = max(ans, a[i+r-1][j+r-1] - a[i-1][j+r-1] - a[i+r-1][j-1] + a[i-1][j-1])
        }
    }
    fmt.Println(ans)
}
