package main
import "fmt"

// implementing big integers manually as it's
// kind of the point of the problem

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
func reverse(s []int) []int {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}
func runeReverse(s []rune) []rune {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}
type BigInt struct {
    number []int
}
func MakeBigInt() BigInt {
    return BigInt{make([]int, 0, 100000)}
}
func (p BigInt) AddInt(q int) BigInt {
    dgts := len(p.number) + 20
    p.number = p.number[:dgts]
    carry := 0
    r := MakeBigInt()
    for i := 0; i < dgts; i++ {
        var current int
        if i == 0 {
            current = p.number[i] + q + carry
        } else {
            current = p.number[i] + carry
        }
        r.number = append(r.number, current % 10)
        carry = current / 10
    }
    for len(r.number) > 0 && r.number[len(r.number)-1] == 0 {
        r.number = r.number[:len(r.number)-1]
    }
    return r
}
func (p BigInt) Add(q BigInt) BigInt {
    dgts := max(len(p.number), len(q.number)) + 1
    p.number = p.number[:dgts]
    q.number = q.number[:dgts]
    carry := 0
    r := MakeBigInt()
    for i := 0; i < dgts; i++ {
        current := p.number[i] + q.number[i] + carry
        r.number = append(r.number, current % 10)
        carry = current / 10
    }
    for len(r.number) > 0 && r.number[len(r.number)-1] == 0 {
        r.number = r.number[:len(r.number)-1]
    }
    return r
}
func (p BigInt) Mult(q int) BigInt {
    dgts := len(p.number) + 20
    p.number = p.number[:dgts]
    carry := 0
    r := MakeBigInt()
    for i := 0; i < dgts; i++ {
        current := p.number[i] * q + carry
        r.number = append(r.number, current % 10)
        carry = current / 10
    }
    for len(r.number) > 0 && r.number[len(r.number)-1] == 0 {
        r.number = r.number[:len(r.number)-1]
    }
    return r
}
func (p BigInt) DivMod(q int) (BigInt, int) {
    dgts := len(p.number)
    p.number = p.number[:dgts]
    carry := 0
    res := make([]int, 0, 100000)
    num := reverse(p.number)
    for i := 0; i < dgts; i++ {
        carry *= 10
        current := (num[i] + carry) / q
        carry = (num[i] + carry) % q
        res = append(res, current)
    }
    r := BigInt{reverse(res)}
    for len(r.number) > 0 && r.number[len(r.number)-1] == 0 {
        r.number = r.number[:len(r.number)-1]
    }
    return r, carry
}
func convert(x rune) int {
    if 'A' <= x && x <= 'Z' {
        return int(x - 'A' + 10)
    } else if 'a' <= x && x <= 'z' {
        return int(x - 'a' + 36)
    } else {
        return int(x - '0')
    }
}
func rconvert(x int) rune {
    if 10 <= x && x <= 35 {
        return rune(x - 10 + 'A')
    } else if 36 <= x && x <= 61 {
        return rune(x - 36 + 'a')
    } else {
        return rune(x + '0')
    }
}
func main() {
    var t int
    fmt.Scanf("%d\n", &t)
    for tt := 0; tt < t; tt++ {
        var from, to int
        var inputs string
        fmt.Scanf("%d %d %s\n", &from, &to, &inputs)
        source := MakeBigInt()
        runes := []rune(inputs)
        // runes = runeReverse(runes)
        for i := 0; i < len(runes); i++ {
            source = source.Mult(from).AddInt(convert(runes[i]))
            // fmt.Println(source)
        }
        output := make([]rune, 0, 100000)
        for len(source.number) > 0 {
            var rem int
            source, rem = source.DivMod(to)
            output = append(output, rconvert(rem))
        }
        output = runeReverse(output)
        strotpt := string(output)
        if len(output) == 0 {
            strotpt = "0"
        }
        fmt.Printf("%d %s\n", from, inputs)
        fmt.Printf("%d %s\n\n", to, strotpt)
    }
}
