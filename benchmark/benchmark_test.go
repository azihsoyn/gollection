package gollection_test

import (
	"sort"
	"testing"

	"github.com/azihsoyn/gollection"
)

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New([]int{})
	}
}

func BenchmarkNew2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New2([]int{})
	}
}

func BenchmarkNewAndResult(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New([]int{}).Result()
	}
}

func BenchmarkNewAndResultAs(b *testing.B) {
	ret := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New([]int{}).ResultAs(&ret)
	}
}

func BenchmarkNewStream(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream([]int{})
	}
}

func BenchmarkDistinct(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New2(arr).Distinct()
	}
}

func BenchmarkDistinct_Stream(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).Distinct().Result()
	}
}

func BenchmarkDistinct_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
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
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).DistinctBy(func(v int) int {
			return v
		}).Result()
	}
}

func BenchmarkDistinctBy_Stream(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).DistinctBy(func(v int) int {
			return v
		}).Result()
	}
}

func BenchmarkDistinctBy_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
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
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).Filter(func(v int) bool {
			return v > 5
		}).Result()
	}
}

func BenchmarkFilter_Stream(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).Filter(func(v int) bool {
			return v > 5
		}).Result()
	}
}

func BenchmarkFilter_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
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
	arr := [][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).FlatMap(func(v int) int {
			return v * 2
		}).Result()
	}
}

func BenchmarkFlatMap_Stream(b *testing.B) {
	arr := [][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).FlatMap(func(v int) int {
			return v * 2
		}).Result()
	}
}

func BenchmarkFlatMap_WithoutGollection(b *testing.B) {
	arr := [][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}}
	b.ResetTimer()
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
	arr := [][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).Flatten().Result()
	}
}

func BenchmarkFlatten_Stream(b *testing.B) {
	arr := [][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).Flatten().Result()
	}
}

func BenchmarkFlatten_WithoutGollection(b *testing.B) {
	arr := [][]int{{0, 0, 1, 1, 2, 2}, {3, 3, 4, 4, 5, 5}, {6, 6, 7, 7, 8, 8, 9, 9}}
	b.ResetTimer()
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
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).Fold(0, func(v1, v2 int) int {
			return v1 + v2
		}).Result()
	}
}

func BenchmarkFold_Stream(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).Fold(0, func(v1, v2 int) int {
			return v1 + v2
		}).Result()
	}
}

func BenchmarkFold_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ret int
		for _, v := range arr {
			ret += v
		}
	}
}

func BenchmarkMap(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).Map(func(v int) int {
			return v * 2
		}).Result()
	}
}

func BenchmarkMap_Stream(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).Map(func(v int) int {
			return v * 2
		}).Result()
	}
}

func BenchmarkMap_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := make([]int, 0, len(arr))
		for _, v := range arr {
			ret = append(ret, v*2)
		}
	}
}

func BenchmarkReduce(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).Reduce(func(v1, v2 int) int {
			return v1 + v2
		}).Result()
	}
}

func BenchmarkReduce_Stream(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).Reduce(func(v1, v2 int) int {
			return v1 + v2
		}).Result()
	}
}

func BenchmarkReduce_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ret int
		for _, v := range arr {
			ret += v
		}
	}
}

func BenchmarkSort(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).SortBy(func(v1, v2 int) bool {
			return v1 < v2
		}).Result()
	}
}

func BenchmarkSort_Stream(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).SortBy(func(v1, v2 int) bool {
			return v1 < v2
		}).Result()
	}
}

func BenchmarkSort_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.IntSlice(arr))
	}
}

func BenchmarkTake(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.New(arr).Take(3).Result()
	}
}

func BenchmarkTake_Stream(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gollection.NewStream(arr).Take(3).Result()
	}
}

func BenchmarkTake_WithoutGollection(b *testing.B) {
	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	b.ResetTimer()
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
