package main
import (
    "fmt"
	"bufio"
	"io/ioutil"
	"os"
)

var bytes []byte
var l int

func fastScan() int {
	b := bytes[l]

	for b < 48 || b > 57 {
		l++
		b = bytes[l]
	}

	result := 0
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int(b-48)

		l++
		if l >= len(bytes) {
			return result
		}
		b = bytes[l]
	}
	return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
    for {
        n := fastScan()
        if n == 0 {
            return
        }
        stk := make([]int, 0, n)
        ind := make([]int, 0, n)
        ans := 0
        for i := 0; i <= n; i++ {
            var a int
            lastind := i
            if i < n {
                a = fastScan()
            }
            for len(stk) > 0 && stk[len(stk)-1] > a {
                ans = max(ans, stk[len(stk)-1] * (i - ind[len(stk)-1]))
                stk = stk[:len(stk)-1]
                lastind = ind[len(ind)-1]
                ind = ind[:len(ind)-1]
            }
            stk = append(stk, a)
            ind = append(ind, lastind)
        }
        fmt.Println(ans)
    }
}
