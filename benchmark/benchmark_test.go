package gollection_test

import (
	"sort"
	"testing"

	"github.com/azihsoyn/gollection"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = gollection.New([]int{})
	}
}

func BenchmarkDistinct(b *testing.B) {
	g := gollection.New([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9})
	for i := 0; i < b.N; i++ {
		g.Distinct()
	}
}

func BenchmarkDistinct_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	for i := 0; i < b.N; i++ {
		m := make(map[int]bool)
		ret := make([]int, 0, len(arr))
		for _, i := range arr {
			if _, ok := m[i]; !ok {
				ret = append(ret, i)
				m[i] = true
			}
		}
	}
}
func BenchmarkDistinctBy(b *testing.B) {
	g := gollection.New([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9})
	for i := 0; i < b.N; i++ {
		g.DistinctBy(func(v int) int {
			return v
		})
	}
}
func BenchmarkDistinctBy_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	f := func(i int) int {
		return i
	}
	for i := 0; i < b.N; i++ {
		m := make(map[int]bool)
		ret := make([]int, 0, len(arr))
		for _, i := range arr {
			id := f(i)
			if _, ok := m[id]; !ok {
				ret = append(ret, i)
				m[id] = true
			}
		}
	}
}

func BenchmarkFilter(b *testing.B) {
	g := gollection.New([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9})
	for i := 0; i < b.N; i++ {
		g.Filter(func(v int) bool {
			return v > 5
		})
	}
}

func BenchmarkFilter_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	for i := 0; i < b.N; i++ {
		ret := make([]int, 0, len(arr))
		for _, i := range arr {
			if i > 5 {
				ret = append(ret, i)
			}
		}
	}
}

func BenchmarkFlatMap(b *testing.B) {
	g := gollection.New([][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}})
	for i := 0; i < b.N; i++ {
		g.FlatMap(func(v int) int {
			return v * 2
		})
	}
}

func BenchmarkFlatMap_WithoutGollection(b *testing.B) {
	arr := [][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}}
	for i := 0; i < b.N; i++ {
		ret := make([]int, 0, len(arr))
		for _, arr2 := range arr {
			for _, v := range arr2 {
				ret = append(ret, v*2)
			}
		}
	}
}

func BenchmarkFlatten(b *testing.B) {
	g := gollection.New([][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}})
	for i := 0; i < b.N; i++ {
		g.Flatten()
	}
}

func BenchmarkFlatten_WithoutGollection(b *testing.B) {
	arr := [][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}}
	for i := 0; i < b.N; i++ {
		ret := make([]int, 0, len(arr))
		for _, arr2 := range arr {
			for _, v := range arr2 {
				ret = append(ret, v)
			}
		}
	}
}

func BenchmarkFold(b *testing.B) {
	g := gollection.New([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9})
	for i := 0; i < b.N; i++ {
		g.Fold(0, func(v1, v2 int) int {
			return v1 + v2
		})
	}
}

func BenchmarkFold_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	for i := 0; i < b.N; i++ {
		var ret int
		for _, v := range arr {
			ret += v
		}
	}
}

func BenchmarkMap(b *testing.B) {
	g := gollection.New([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9})
	for i := 0; i < b.N; i++ {
		g.Map(func(v int) int {
			return v * 2
		})
	}
}

func BenchmarkMap_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	for i := 0; i < b.N; i++ {
		ret := make([]int, 0, len(arr))
		for _, v := range arr {
			ret = append(ret, v*2)
		}
	}
}
func BenchmarkReduce(b *testing.B) {
	g := gollection.New([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9})
	for i := 0; i < b.N; i++ {
		g.Reduce(func(v1, v2 int) int {
			return v1 + v2
		})
	}
}

func BenchmarkReduce_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	for i := 0; i < b.N; i++ {
		var ret int
		for _, v := range arr {
			ret += v
		}
	}
}

func BenchmarkSort(b *testing.B) {
	g := gollection.New([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9})
	for i := 0; i < b.N; i++ {
		g.SortBy(func(v1, v2 int) bool {
			return v1 < v2
		})
	}
}

func BenchmarkSort_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.IntSlice(arr))
	}
}

func BenchmarkTake(b *testing.B) {
	g := gollection.New([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9})
	for i := 0; i < b.N; i++ {
		g.Take(3)
	}
}

func BenchmarkTake_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	for i := 0; i < b.N; i++ {
		limit := 3
		if limit < len(arr) {
			limit = len(arr)
		}
		ret := make([]int, 0, limit)
		for i := 0; i < limit; i++ {
			ret = append(ret, arr[i])
		}
	}
}
