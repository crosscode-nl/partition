package partition

// ToChan returns a chan that receives indices for all parts.
func ToChan(totalLength int, partitionLength int) <-chan Indices {
	c := make(chan Indices)
	go func() {
		ToFunc(totalLength, partitionLength, func(l int, h int) {
			c <- Indices{Low: l, High: h}
		})
		close(c)
	}()
	return c
}
