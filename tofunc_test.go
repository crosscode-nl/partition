package partition_test

import (
	"fmt"
	"gitlab.crosscode.nl/metamsglabs/metamsg/partition"
	"reflect"
	"testing"
)

func TestToFunc(t *testing.T) {
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
			var got []partition.Indices
			partition.ToFunc(tt.args.totalLength, tt.args.partitionLength, func(low int, high int) {
				got = append(got, partition.Indices{Low: low, High: high})
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleToFunc() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	partition.ToFunc(len(a), 5, func(l int, h int) {
		fmt.Printf("Part: %v\n", a[l:h])
	})
	// Output:
	// Part: [1 2 3 4 5]
	// Part: [6 7 8 9]
}

func BenchmarkToFunc100(b *testing.B) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sum int

	for i := 0; i < b.N; i++ {
		partition.ToFunc(len(a), 5, func(l int, h int) {
			for _, p := range a[l:h] {
				sum += p
			}
		})
	}
}

func BenchmarkToFunc10(b *testing.B) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sum int

	for i := 0; i < b.N; i++ {
		partition.ToFunc(len(a), 5, func(l int, h int) {
			for _, p := range a[l:h] {
				sum += p
			}
		})
	}
}

func BenchmarkToFunc0(b *testing.B) {
	var a []int
	var sum int

	for i := 0; i < b.N; i++ {
		partition.ToFunc(len(a), 5, func(l int, h int) {
			for _, p := range a[l:h] {
				sum += p
			}
		})
	}
}
