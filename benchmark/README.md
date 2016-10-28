# Benchmark

```
BenchmarkNew-8                                20000000          78.9 ns/op           0 B/op           0 allocs/op
BenchmarkDistinct-8                             100000         15453 ns/op        1123 B/op          44 allocs/op
BenchmarkDistinct_WithoutGollection-8           500000          3342 ns/op         339 B/op           2 allocs/op
BenchmarkDistinctBy-8                            30000         43210 ns/op        1843 B/op          54 allocs/op
BenchmarkDistinctBy_WithoutGollection-8         500000          3529 ns/op         339 B/op           2 allocs/op
BenchmarkFilter-8                                50000         37765 ns/op        1456 B/op          51 allocs/op
BenchmarkFilter_WithoutGollection-8            3000000           541 ns/op         160 B/op           1 allocs/op
BenchmarkFlatMap-8                               30000         45710 ns/op        2144 B/op          69 allocs/op
BenchmarkFlatMap_WithoutGollection-8           1000000          1379 ns/op         368 B/op           4 allocs/op
BenchmarkFlatten-8                              100000         14287 ns/op        1184 B/op          29 allocs/op
BenchmarkFlatten_WithoutGollection-8           1000000          1455 ns/op         368 B/op           4 allocs/op
BenchmarkFold-8                                  30000         41176 ns/op        1336 B/op          42 allocs/op
BenchmarkFold_WithoutGollection-8             10000000           228 ns/op           0 B/op           0 allocs/op
BenchmarkMap-8                                   30000         42234 ns/op        1840 B/op          63 allocs/op
BenchmarkMap_WithoutGollection-8               2000000           714 ns/op         160 B/op           1 allocs/op
BenchmarkReduce-8                                50000         38513 ns/op        1272 B/op          40 allocs/op
BenchmarkReduce_WithoutGollection-8           10000000           226 ns/op           0 B/op           0 allocs/op
BenchmarkSort-8                                  10000        137307 ns/op        4416 B/op         129 allocs/op
BenchmarkSort_WithoutGollection-8               300000          4635 ns/op          32 B/op           1 allocs/op
BenchmarkTake-8                                1000000          2515 ns/op         336 B/op           6 allocs/op
BenchmarkTake_WithoutGollection-8              2000000           706 ns/op         160 B/op           1 allocs/op
PASS
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
