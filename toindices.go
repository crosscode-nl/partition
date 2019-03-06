package partition

// ToIndices returns []Indices
func ToIndices(totalLength int, partitionLength int) []Indices {
	var results []Indices
	ToFunc(totalLength, partitionLength, func(l int, h int) {
		results = append(results, Indices{Low: l, High: h})
	})
	return results
}
