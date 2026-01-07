package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	values := []int{}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
		values = append(values, a[i])
	}

	fmt.Scanf("%d", &m)
	b := make([]int, m)
	c := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
		values = append(values, b[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &c[i])
		values = append(values, c[i])
	}

	// 离散化
	sort.Ints(values)
	uniq := []int{}
	for _, v := range values {
		if len(uniq) == 0 || uniq[len(uniq)-1] != v {
			uniq = append(uniq, v)
		}
	}

	rank := make(map[int]int)
	for i, v := range uniq {
		rank[v] = i // 映射原值到离散化后的索引
	}

	// 用离散化索引统计数量
	speakers := make([]int, len(uniq))
	for i := 0; i < n; i++ {
		speakers[rank[a[i]]]++
	}

	// 找最佳 id
	bestNative, bestTrans, id := -1, -1, 0
	for i := 0; i < m; i++ {
		curnat := speakers[rank[b[i]]]
		curtrans := speakers[rank[c[i]]]
		if curnat > bestNative || (curnat == bestNative && curtrans > bestTrans) {
			bestNative = curnat
			bestTrans = curtrans
			id = i + 1 // 题目是从 1 开始编号
		}
	}

	fmt.Println(id)
}

