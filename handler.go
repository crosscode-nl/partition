// Package partition contains methods to help partition slices.
// Benchmark results show that the ToFunc function is the fastest option available in this package.
// BenchmarkToChan100-4      	  300000	      4166 ns/op
// BenchmarkToChan10-4       	 2000000	       840 ns/op
// BenchmarkToChan0-4        	 3000000	       435 ns/op
// BenchmarkToFunc100-4      	10000000	       197 ns/op
// BenchmarkToFunc10-4       	50000000	        24.9 ns/op
// BenchmarkToFunc0-4        	500000000	         3.05 ns/op
// BenchmarkToHandler100-4   	 3000000	       590 ns/op
// BenchmarkToHandler10-4    	10000000	       127 ns/op
// BenchmarkToHandler0-4     	50000000	        38.8 ns/op
// BenchmarkToIndices10-4    	20000000	        90.0 ns/op
// BenchmarkToIndices100-4   	 3000000	       504 ns/op
// BenchmarkToIndices0-4     	300000000	         4.98 ns/op
package partition

// Handler interface defines an interface for a partition handler that Partition accepts.
type Handler interface {
	Handle(lowIndex int, highIndex int)
}
