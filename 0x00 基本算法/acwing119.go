package main
import (
    "fmt"
    "sort"
    "math"
    "bufio"
	"io/ioutil"
	"os"
)

var mindist float64

type Point struct {
    x int
    y int
    tp bool
}

func update(a Point, b Point) {
    if a.tp != b.tp {
        dist :=
          math.Sqrt(float64((a.x - b.x) * (a.x - b.x) + (a.y - b.y) * (a.y - b.y)))
        if dist < mindist {
            mindist = dist
        }
    }
}

var pts []Point

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

func conquer(l, r int) {
    if l == r {
        return
    }
    m := (l + r) / 2
    midx := pts[m].x
    conquer(l, m)
    conquer(m + 1, r)
    ll := l
    rr := m + 1
    tmp := make([]Point, 0, r - l + 1)
    for i := l; i <= r; i++ {
        if rr > r || ll <= m && pts[ll].y < pts[rr].y {
            tmp = append(tmp, pts[ll])
            ll += 1
        } else {
            tmp = append(tmp, pts[rr])
            rr += 1
        }
    }
    tmp2 := make([]Point, 0, r - l + 1)
    for i := 0; i <= r - l; i++ {
        pts[l + i] = tmp[i]
        if float64(midx) - mindist < float64(pts[l+i].x) && float64(pts[l+i].x) < float64(midx) + mindist {
            tmp2 = append(tmp2, pts[l+i])
        }
    }
    // fmt.Println(mindist)
    // fmt.Println(l, r, tmp2)
    for i := 0; i < len(tmp2); i++ {
        // 不要枚举太多，要不然肯能会被卡。
        // 标称不卡的方式是直接不把它放到另一个数组里，
        // 这种写法可以通过。不能按照标准的方法写，因为
        // 那种方法依赖于 mindist 更新，但是有些数据
        // 的时候不一定都可以更新。
        for j := i - 1; j >= 0 && j >= i - 7; j-- {
            update(tmp2[i], tmp2[j])
        }
    }
}

func main() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
    t := fastScan()
    for tt := 0; tt < t; tt++ {
        n := fastScan()
        pts = make([]Point, 2 * n)
        for i := 0; i < n; i++ {
            pts[i].x, pts[i].y = fastScan(), fastScan()
        }
        for i := n; i < 2 * n; i++ {
            pts[i].x, pts[i].y = fastScan(), fastScan()
            pts[i].tp = true
        }
    	sort.Slice(pts, func(i, j int) bool { return pts[i].x < pts[j].x })
        // fmt.Println(pts)
    	mindist = 2e9
    	conquer(0, len(pts) - 1)
    	fmt.Printf("%.3f\n", mindist)
    }
}

