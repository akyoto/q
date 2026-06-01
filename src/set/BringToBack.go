package set

// BringToBack brings the element at `index` within the `slice` to the back.
func BringToBack[T any](slice []T, index int) {
	if index < 0 || index >= len(slice)-1 {
		return
	}

	target := slice[index]

	for i := index; i < len(slice)-1; i++ {
		slice[i] = slice[i+1]
	}

	slice[len(slice)-1] = target
}