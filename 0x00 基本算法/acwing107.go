package main
import (
    "fmt"
)

func merge(ls []int) int {
    if len(ls) == 1 {
        return 0
    }
    mid := len(ls)/2
    left_pairs := merge(ls[:mid])
    right_pairs := merge(ls[mid:])
    merged := make([]int, 0, len(ls))
    i := 0
    j := mid
    cnt := 0
    for k := 0; k < len(ls); k++ {
        if j >= len(ls) || i <= mid - 1 && ls[i] < ls[j] {
            merged = append(merged, ls[i])
            i += 1
        } else {
            merged = append(merged, ls[j])
            j += 1
            cnt += mid - i
        }
    }
    for k := 0; k < len(ls); k++ {
        ls[k] = merged[k]
    }
    return cnt + left_pairs + right_pairs
}

func main() {
    var n int
    for {
        fmt.Scanf("%d", &n)
        if n == 0 {
            break
        }
        a := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Scanf("%d", &a[i])
        }
        fmt.Println(merge(a))
    }
}
