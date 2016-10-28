# Benchmark

```
BenchmarkNew-8                                20000000          89.4 ns/op           0 B/op           0 allocs/op
BenchmarkDistinct-8                             100000         14808 ns/op        1043 B/op          34 allocs/op
BenchmarkDistinct_WithoutGollection-8           500000          3534 ns/op         339 B/op           2 allocs/op
BenchmarkDistinctBy-8                            30000         49867 ns/op        1843 B/op          54 allocs/op
BenchmarkDistinctBy_WithoutGollection-8         300000          3928 ns/op         339 B/op           2 allocs/op
BenchmarkFilter-8                                30000         41066 ns/op        1456 B/op          51 allocs/op
BenchmarkFilter_WithoutGollection-8            3000000           551 ns/op         160 B/op           1 allocs/op
BenchmarkFlatMap-8                               30000         51068 ns/op        2144 B/op          69 allocs/op
BenchmarkFlatMap_WithoutGollection-8           1000000          1393 ns/op         368 B/op           4 allocs/op
BenchmarkFlatten-8                              100000         15726 ns/op        1184 B/op          29 allocs/op
BenchmarkFlatten_WithoutGollection-8           1000000          1388 ns/op         368 B/op           4 allocs/op
BenchmarkFold-8                                  30000         45872 ns/op        1336 B/op          42 allocs/op
BenchmarkFold_WithoutGollection-8              5000000           254 ns/op           0 B/op           0 allocs/op
BenchmarkMap-8                                   30000         46036 ns/op        1840 B/op          63 allocs/op
BenchmarkMap_WithoutGollection-8               2000000           702 ns/op         160 B/op           1 allocs/op
BenchmarkReduce-8                                30000         39491 ns/op        1272 B/op          40 allocs/op
BenchmarkReduce_WithoutGollection-8            5000000           240 ns/op           0 B/op           0 allocs/op
BenchmarkSort-8                                  10000        148140 ns/op        4224 B/op         127 allocs/op
BenchmarkSort_WithoutGollection-8               300000          5172 ns/op          32 B/op           1 allocs/op
BenchmarkTake-8                                1000000          2483 ns/op         336 B/op           6 allocs/op
BenchmarkTake_WithoutGollection-8              2000000           701 ns/op         160 B/op           1 allocs/op
```

machine spec
```
$ system_profiler SPHardwareDataType
Hardware:

    Hardware Overview:

      Model Name: MacBook Pro
      Model Identifier: MacBookPro11,4
      Processor Name: Intel Core i7
      Processor Speed: 2.2 GHz
      Number of Processors: 1
      Total Number of Cores: 4
      L2 Cache (per Core): 256 KB
      L3 Cache: 6 MB
      Memory: 16 GB
```
