main() {
	t0 := -50
	t1 := -8
	t2 := -50
	t3 := -29
	s := 0

	loop i := 0..1 {
		s = s + t0
	}

	bb := 0

	switch {
		t0 > 0 { bb = 10 }
		_      { bb = -46 }
	}

	assert (t1 << 0) - bb == 38
}