# partition [![Go Report Card](https://goreportcard.com/badge/github.com/crosscode-nl/partition)](https://goreportcard.com/report/github.com/crosscode-nl/partition)
GO/golang package for partitioning slices, arrays and strings. 

## Introduction

I needed a solution for partitioning a slice. I searched for an existing package, and found the following StackOverflow question: https://stackoverflow.com/questions/30242567/in-go-how-can-i-partition-a-slice-array-string-of-any-type

This pointed to a package which only has an implementation that uses chans. That has too much overhead to my liking so I wrote this package, but added more flexibility in how to use this package, each with its own performance or memory usage characteristics. 

This package is fully unit tested and documented.

### Usage

Download the package with the following command:

~~~console
go get github.com/crosscode-nl/partition
~~~

Use the following import in your code to use the package.

~~~go
import "github.com/crosscode-nl/partition"
~~~

## Performance

The following benchmarks show that the partition.ToFunc function is the fastest to use. 
~~~
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
~~~

## Examples

### partition.ToFunc

This func takes a func that receives all partitions one by one.

~~~go
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	partition.ToFunc(len(a), 5, func(l int, h int) {
		fmt.Printf("Part: %v\n", a[l:h])
	})
	// Output:
	// Part: [1 2 3 4 5]
	// Part: [6 7 8 9]
~~~

### partition.ToChan

This func returns a chan that receives all partitions one by one.

~~~go
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for part := range partition.ToChan(len(a), 5) {
		fmt.Printf("Part: %v\n", a[part.Low:part.High])
	}
	// Output:
	// Part: [1 2 3 4 5]
	// Part: [6 7 8 9]
~~~

### partition.ToHandler

This func takes an implementation of the partition.Handler interface.  

~~~go
// Handler interface defines an interface for a partition handler that Partition accepts.
type Handler interface {
	Handle(lowIndex int, highIndex int)
}
~~~ 

The example contains an implementation used by the unit tests partition.SliceHandler

~~~go
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := partition.SliceHandler{}
	partition.ToHandler(len(a), 5, &s)
	for _, part := range s.Parts {
		fmt.Printf("Part: %v\n", a[part.Low:part.High])
	}
	// Output:
	// Part: [1 2 3 4 5]
	// Part: [6 7 8 9]
~~~

### partition.ToIndices

This func returns a slice with all partitions.

~~~go
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, part := range partition.ToIndices(len(a), 5) {
		fmt.Printf("Part: %v\n", a[part.Low:part.High])
	}
	// Output:
	// Part: [1 2 3 4 5]
	// Part: [6 7 8 9]
~~~

## License

MIT License

Copyright (c) 2019 CrossCode / P. Vollebregt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
