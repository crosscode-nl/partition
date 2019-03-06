# partition

GO/golang package for partitioning slices, arrays and strings. 

## Introduction

I needed a solution for partitioning a slice. I searched for an existing package, and found the following StackOverflow question: https://stackoverflow.com/questions/30242567/in-go-how-can-i-partition-a-slice-array-string-of-any-type

This pointed to a package which only has an implementation that uses chans. That has to much overhead too my liking so I wrote this package, but added more options of using this package to add flexibility of using this package.

## Performance

The following benchmarks show that the partition.ToFunc function is the fastest to use. 

BenchmarkToChan100-4      	   50000	     25862 ns/op
BenchmarkToChan10-4       	  300000	      4939 ns/op
BenchmarkToChan0-4        	  500000	      2727 ns/op
BenchmarkToFunc100-4      	 5000000	       230 ns/op
BenchmarkToFunc10-4       	30000000	        46.5 ns/op
BenchmarkToFunc0-4        	100000000	        14.6 ns/op
BenchmarkToHandler100-4   	  500000	      3877 ns/op
BenchmarkToHandler10-4    	 1000000	      1233 ns/op
BenchmarkToHandler0-4     	 3000000	       478 ns/op
BenchmarkToIndices10-4    	 2000000	       739 ns/op
BenchmarkToIndices100-4   	  500000	      3241 ns/op
BenchmarkToIndices0-4     	50000000	        26.5 ns/op

## Examples

### ToFunc

~~~.go
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	partition.ToFunc(len(a), 5, func(l int, h int) {
		fmt.Printf("Part: %v\n", a[l:h])
	})
	// Output:
	// Part: [1 2 3 4 5]
	// Part: [6 7 8 9]
~~~
