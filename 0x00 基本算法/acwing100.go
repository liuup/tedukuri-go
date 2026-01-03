package main
import "fmt"
func main() {
    var n int
    var last int
    fmt.Scanf("%d\n%d", &n, &last)
    possum := 0
    negsum := 0
    for i := 2; i <= n; i++ {
        var cur int
        fmt.Scanf("%d", &cur)
        if cur - last > 0 {
            possum += cur - last
        } else {
            negsum += last - cur
        }
        last = cur
    }
    if possum > negsum {
        fmt.Println(possum)
        fmt.Println(possum - negsum + 1)
    } else {
        fmt.Println(negsum)
        fmt.Println(negsum - possum + 1)
    }
}
