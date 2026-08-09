sort(x []int) {
	quicksort(x, 0, x.len-1 as int)
}

quicksort(x []int, low int, high int) {
	loop {
		if low >= high {
			return
		}

		i := partition(x, low, high)

		if i - low < high - i {
			quicksort(x, low, i-1)
			low = i + 1
		} else {
			quicksort(x, i+1, high)
			high = i - 1
		}
	}
}

partition(x []int, low int, high int) -> int {
	pivot := x[high]
	i := low
	j := low

	loop {
		if j >= high {
			swap(x, i, high)
			return i
		}

		if x[j] < pivot {
			swap(x, i, j)
			i += 1
		}

		j += 1
	}
}

swap(x []int, i int, j int) {
	tmp := x[i]
	x[i] = x[j]
	x[j] = tmp
}