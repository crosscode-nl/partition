package partition

// HandleFunc type defines the func that is used as a callback in ToFunc
type HandleFunc func(low int, high int)

// ToFunc calls provided func per fragment
func ToFunc(totalLength int, partitionLength int, hf HandleFunc) {
	if partitionLength <= 0 || totalLength <= 0 {
		return
	}
	partitions := totalLength / partitionLength
	var i int
	for i = 0; i < partitions; i++ {
		hf(i*partitionLength, i*partitionLength+partitionLength)
	}
	if rest := totalLength % partitionLength; rest != 0 {
		hf(i*partitionLength, i*partitionLength+rest)
	}
}
