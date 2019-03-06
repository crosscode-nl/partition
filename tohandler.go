package partition

// ToHandler calls the handler for each part.
func ToHandler(totalLength int, partitionLength int, h Handler) {
	ToFunc(totalLength, partitionLength, h.Handle)
}
