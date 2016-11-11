# Benchmark

```
$ go test -bench . -benchmem
BenchmarkNew-8                            	20000000	        99.5 ns/op	      96 B/op	       2 allocs/op
BenchmarkNewStream-8                      	  200000	     10247 ns/op	     289 B/op	       8 allocs/op
BenchmarkDistinct-8                       	  300000	      4267 ns/op	    1155 B/op	      36 allocs/op
BenchmarkDistinct_Stream-8                	    5000	    265562 ns/op	    1717 B/op	      49 allocs/op
BenchmarkDistinct_WithoutGollection-8     	 1000000	      1211 ns/op	     339 B/op	       2 allocs/op
BenchmarkDistinctBy-8                     	  200000	      9064 ns/op	    1955 B/op	      56 allocs/op
BenchmarkDistinctBy_Stream-8              	    5000	    296762 ns/op	    2678 B/op	      89 allocs/op
BenchmarkDistinctBy_WithoutGollection-8   	 1000000	      1267 ns/op	     339 B/op	       2 allocs/op
BenchmarkFilter-8                         	  200000	      7145 ns/op	    1568 B/op	      53 allocs/op
BenchmarkFilter_Stream-8                  	   10000	    157181 ns/op	    1952 B/op	      83 allocs/op
BenchmarkFilter_WithoutGollection-8       	20000000	        84.8 ns/op	     160 B/op	       1 allocs/op
BenchmarkFlatMap-8                        	  200000	      8498 ns/op	    2256 B/op	      71 allocs/op
BenchmarkFlatMap_Stream-8                 	   10000	    144181 ns/op	    2656 B/op	      80 allocs/op
BenchmarkFlatMap_WithoutGollection-8      	 5000000	       240 ns/op	     368 B/op	       4 allocs/op
BenchmarkFlatten-8                        	  500000	      2990 ns/op	    1296 B/op	      31 allocs/op
BenchmarkFlatten_Stream-8                 	   10000	    121338 ns/op	    1856 B/op	      60 allocs/op
BenchmarkFlatten_WithoutGollection-8      	10000000	       232 ns/op	     368 B/op	       4 allocs/op
BenchmarkFold-8                           	  200000	      6847 ns/op	    1448 B/op	      44 allocs/op
BenchmarkFold_Stream-8                    	   10000	    100037 ns/op	    1744 B/op	      68 allocs/op
BenchmarkFold_WithoutGollection-8         	100000000	        10.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap-8                            	  200000	      7645 ns/op	    1952 B/op	      65 allocs/op
BenchmarkMap_Stream-8                     	   10000	    135785 ns/op	    2720 B/op	      97 allocs/op
BenchmarkMap_WithoutGollection-8          	20000000	        83.7 ns/op	     160 B/op	       1 allocs/op
BenchmarkReduce-8                         	  200000	      7254 ns/op	    1384 B/op	      42 allocs/op
BenchmarkReduce_Stream-8                  	   10000	    113624 ns/op	    1680 B/op	      65 allocs/op
BenchmarkReduce_WithoutGollection-8       	100000000	        14.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSort-8                           	  100000	     22518 ns/op	    4336 B/op	     129 allocs/op
BenchmarkSort_Stream-8                    	    5000	    222876 ns/op	    7136 B/op	     232 allocs/op
BenchmarkSort_WithoutGollection-8         	 3000000	       437 ns/op	      32 B/op	       1 allocs/op
BenchmarkTake-8                           	 3000000	       600 ns/op	     448 B/op	       8 allocs/op
BenchmarkTake_Stream-8                    	   50000	     34518 ns/op	    1221 B/op	      24 allocs/op
BenchmarkTake_WithoutGollection-8         	 5000000	       229 ns/op	     160 B/op	       1 allocs/op
PASS
ok  	github.com/azihsoyn/gollection/benchmark	51.777s
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
