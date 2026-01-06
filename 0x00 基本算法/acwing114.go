package main
import (
    "fmt"
    "math/big"
    "sort"
)
type Slave struct {
    a int64
    b int64
}
func biMax(a, b *big.Int) *big.Int {
	if a.Cmp(b) >= 0 {
		return a
	}
	return b
}
func main() {
    var n int
    var a, b int64
    fmt.Scanf("%d", &n)
    fmt.Scanf("%d %d", &a, &b)
    slaves := make([]*Slave, n)
    for i := 0; i < n; i++ {
        slaves[i] = &Slave{}
        fmt.Scanf("%d %d", &slaves[i].a, &slaves[i].b)
    }
    sort.Slice(slaves, func (i, j int) bool {
        return slaves[i].a * slaves[i].b < slaves[j].a * slaves[j].b
    })
    prod := &big.Int{}
    prod.SetInt64(a)
    ans := &big.Int{}
    for i := 0; i < n; i++ {
        reg := &big.Int{}
        reg.SetInt64(slaves[i].b)
        reg2 := &big.Int{}
        reg2.Div(prod, reg)
        ans = biMax(ans, reg2)
        reg.SetInt64(slaves[i].a)
        prod.Mul(prod, reg)
    }
    fmt.Println(ans)
}
