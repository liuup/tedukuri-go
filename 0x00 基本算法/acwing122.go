package main
import (
    "fmt"
    "sort"
    "io/ioutil"
    "bufio"
    "os"
)
func iAbs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}
var bytes []byte
var l, max int

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
		if l > max {
			return result
		}
		b = bytes[l]
	}
	return result
}
func main() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
    n := fastScan()
    a := make([]int, n)
    avg := 0
    for i := 0; i < n; i++ {
        a[i] = fastScan()
        avg += a[i]
    }
    avg /= n
    for i := 0; i < n; i++ {
        if i > 0 {
            a[i] += a[i-1]
        }
        a[i] = a[i] - avg
    }
    sort.Ints(a)
    me := a[n/2]
    ans := 0
    for i := 0; i < n; i++ {
        ans += iAbs(me - a[i])
    }
    fmt.Println(ans)
}
