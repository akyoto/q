// Copyright (C) 2016 Felipe da Cunha Gonçalves
// Source: https://github.com/xrash/smetrics
// Algorithm: William E. Winkler (1990) and Matthew A. Jaro (1989)
package similarity

import (
	"math"
)

// Default compares two strings and returns the similarity between 0 and 1.
func Default(a string, b string) float64 {
	return JaroWinkler(a, b, 0.7, 4)
}

func JaroWinkler(a, b string, boostThreshold float64, prefixSize int) float64 {
	j := Jaro(a, b)

	if j <= boostThreshold {
		return j
	}

	prefixSize = int(math.Min(float64(len(a)), math.Min(float64(prefixSize), float64(len(b)))))
	var prefixMatch float64

	for i := 0; i < prefixSize; i++ {
		if a[i] == b[i] {
			prefixMatch++
		}
	}

	return j + 0.1*prefixMatch*(1.0-j)
}

func Jaro(a, b string) float64 {
	la := float64(len(a))
	lb := float64(len(b))

	matchRange := int(math.Floor(math.Max(la, lb)/2.0)) - 1
	matchRange = int(math.Max(0, float64(matchRange-1)))
	var matches, halfs float64
	transposed := make([]bool, len(b))

	for i := 0; i < len(a); i++ {
		start := int(math.Max(0, float64(i-matchRange)))
		end := int(math.Min(lb-1, float64(i+matchRange)))

		for j := start; j <= end; j++ {
			if transposed[j] {
				continue
			}

			if a[i] == b[j] {
				if i != j {
					halfs++
				}

				matches++
				transposed[j] = true
				break
			}
		}
	}

	if matches == 0 {
		return 0
	}

	transposes := math.Floor(float64(halfs / 2))
	return ((matches / la) + (matches / lb) + (matches-transposes)/matches) / 3.0
}
