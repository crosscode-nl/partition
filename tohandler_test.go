package partition_test

import (
	"fmt"
	"github.com/crosscode-nl/partition"
	"reflect"
	"testing"
)

func TestToHandler(t *testing.T) {
	type args struct {
		totalLength     int
		partitionLength int
	}
	tests := []struct {
		name string
		args args
		want []partition.Indices
	}{
		{name: "l1p1", args: args{totalLength: 1, partitionLength: 1}, want: []partition.Indices{{Low: 0, High: 1}}},
		{name: "l2p1", args: args{totalLength: 2, partitionLength: 1}, want: []partition.Indices{{Low: 0, High: 1}, {Low: 1, High: 2}}},
		{name: "l2p2", args: args{totalLength: 2, partitionLength: 2}, want: []partition.Indices{{Low: 0, High: 2}}},
		{name: "l2p3", args: args{totalLength: 2, partitionLength: 3}, want: []partition.Indices{{Low: 0, High: 2}}},
		{name: "l6p3", args: args{totalLength: 6, partitionLength: 3}, want: []partition.Indices{{Low: 0, High: 3}, {Low: 3, High: 6}}},
		{name: "l6p5", args: args{totalLength: 6, partitionLength: 5}, want: []partition.Indices{{Low: 0, High: 5}, {Low: 5, High: 6}}},
		{name: "l6p0", args: args{totalLength: 6, partitionLength: 0}, want: nil},
		{name: "l0p6", args: args{totalLength: 0, partitionLength: 6}, want: nil},
		{name: "l0p0", args: args{totalLength: 0, partitionLength: 0}, want: nil},
		{name: "l-1p6", args: args{totalLength: -1, partitionLength: 6}, want: nil},
		{name: "l6p-1", args: args{totalLength: 6, partitionLength: -1}, want: nil},
		{name: "l-1p-1", args: args{totalLength: -1, partitionLength: -1}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := partition.SliceHandler{}
			partition.ToHandler(tt.args.totalLength, tt.args.partitionLength, &s)
			got := s.Parts
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleToHandler() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := partition.SliceHandler{}
	partition.ToHandler(len(a), 5, &s)
	for _, part := range s.Parts {
		fmt.Printf("Part: %v\n", a[part.Low:part.High])
	}
	// Output:
	// Part: [1 2 3 4 5]
	// Part: [6 7 8 9]
}

func BenchmarkToHandler100(b *testing.B) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sum int

	for i := 0; i < b.N; i++ {
		s := partition.SliceHandler{}
		partition.ToHandler(len(a), 5, &s)
		for _, part := range s.Parts {
			for _, p := range a[part.Low:part.High] {
				sum += p
			}
		}
	}
}

func BenchmarkToHandler10(b *testing.B) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sum int

	for i := 0; i < b.N; i++ {
		s := partition.SliceHandler{}
		partition.ToHandler(len(a), 5, &s)
		for _, part := range s.Parts {
			for _, p := range a[part.Low:part.High] {
				sum += p
			}
		}
	}
}

func BenchmarkToHandler0(b *testing.B) {
	var a []int
	var sum int

	for i := 0; i < b.N; i++ {
		s := partition.SliceHandler{}
		partition.ToHandler(len(a), 5, &s)
		for _, part := range s.Parts {
			for _, p := range a[part.Low:part.High] {
				sum += p
			}
		}
	}
}
