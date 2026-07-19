main() {
	c := 0

	switch 42 {
		41 { c -= 1 }
	}

	switch 42 {
		41 { c -= 1 }
		_  { c += 1 }
	}

	switch 42 {
		41 { c -= 1 }
		42 { c += 1 }
		_  { c -= 1 }
	}

	switch "b" {
		"a" { c -= 1 }
	}

	switch "b" {
		"a" { c -= 1 }
		_   { c += 1 }
	}

	switch "b" {
		"a" { c -= 1 }
		"b" { c += 1 }
		_   { c -= 1 }
	}

	assert c == 4
}