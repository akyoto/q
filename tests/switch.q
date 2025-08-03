main() {
	c := 0

	switch {
		1 == 1 { c += 1 }
	}

	switch {
		1 == 1 { c += 1 }
		_      { c -= 1 }
	}

	switch {
		0 == 1 { c -= 1 }
		_      { c += 1 }
	}

	switch {
		0 == 1 { c -= 1 }
		0 == 2 { c -= 1 }
		_      { c += 1 }
	}

	switch {
		0 == 1 { c -= 1 }
		0 == 2 { c -= 1 }
		2 == 2 { c += 1 }
		_      { c -= 1 }
	}

	switch {
		0 == 1 { c -= 1 }
		0 == 2 { c -= 1 }
		0 == 3 { c -= 1 }
	}

	assert c == 5
}