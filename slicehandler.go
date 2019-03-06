package partition

// SliceHandler is an implementation of partition.Handler and is a handler that is used internally and for tests and examples.
type SliceHandler struct {
	Parts []Indices
}

// Handle implements handling a partition part.
func (h *SliceHandler) Handle(lowIndex int, highIndex int) {
	h.Parts = append(h.Parts, Indices{Low: lowIndex, High: highIndex})
}
