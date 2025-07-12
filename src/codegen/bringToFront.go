package codegen

// bringToFront brings the element at `index` within the `slice` to the front.
func bringToFront[T any](slice []T, index int) {
	if index <= 0 || index >= len(slice) {
		return
	}

	target := slice[index]

	for i := index; i > 0; i-- {
		slice[i] = slice[i-1]
	}

	slice[0] = target
}