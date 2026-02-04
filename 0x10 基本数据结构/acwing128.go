package main
import (
    "fmt"
	"io/ioutil"
	"os"
	"bufio"
)
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
var bytes []byte
var l int

func fastScan() int {

    for l < len(bytes) && (bytes[l] == ' ' || bytes[l] == '\t' || bytes[l] == '\n' || bytes[l] == '\r') {
        l++
    }
    
    sign := 1
    if l < len(bytes) && bytes[l] == '-' {
        sign = -1
        l++
    }

	b := bytes[l]

	result := 0
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int(b-48)

		l++
		if l > len(bytes)-1 {
			return result * sign
		}
		b = bytes[l]
	}
	return result * sign
}
func fastScan1() rune {
    for l < len(bytes) && (bytes[l] == ' ' || bytes[l] == '\t' || bytes[l] == '\n' || bytes[l] == '\r') {
        l++
    }
	b := bytes[l]
	l++
	return rune(b)
}
func main() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))

    var q int
    q = fastScan()
    stk1 := make([]int, 0, 1000000)
    sum1 := make([]int, 0, 1000000)
    max1 := make([]int, 0, 1000000)
    stk2 := make([]int, 0, 1000000)
    for qq := 1; qq <= q; qq++ {
        var y rune
        y = fastScan1()
        if y == 'I' {
            var x int
            x = fastScan()
            stk1 = append(stk1, x)
            if len(sum1) > 0 {
                sum1 = append(sum1, x+sum1[len(sum1)-1])
                max1 = append(max1, max(max1[len(max1)-1], sum1[len(sum1)-1]))
            } else {
                sum1 = append(sum1, x)
                max1 = append(max1, x)
            }
        } else if y == 'D' {
            if len(stk1) > 0 {
                stk1 = stk1[:len(stk1)-1]
                sum1 = sum1[:len(sum1)-1]
                max1 = max1[:len(max1)-1]
            }
        } else if y == 'L' {
            if len(stk1) > 0 {
                stk2 = append(stk2, stk1[len(stk1)-1])
                stk1 = stk1[:len(stk1)-1]
                sum1 = sum1[:len(sum1)-1]
                max1 = max1[:len(max1)-1]
            }
        } else if y == 'R' {
            if len(stk2) > 0 {
                x := stk2[len(stk2)-1]
                stk2 = stk2[:len(stk2)-1]
                stk1 = append(stk1, x)
                if len(sum1) > 0 {
                    sum1 = append(sum1, x+sum1[len(sum1)-1])
                    max1 = append(max1, max(max1[len(max1)-1], sum1[len(sum1)-1]))
                } else {
                    sum1 = append(sum1, x)
                    max1 = append(max1, x)
                }
            }
        } else if y == 'Q' {
            var k int
            k = fastScan()
            // fmt.Println(stk1, sum1, max1, stk2, k)
            fmt.Println(max1[k-1])
        } else {
            fmt.Println("Assert")
        }
    }
}
