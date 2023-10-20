package main

import (
	// "fmt"
	"fmt"
	"math/big"

	fft "github.com/Rusih100/polynomial"
)

var alphabet []string

func init() {
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}
}

func CompareFiles(precision float64, first, second FileAndContent) { // assume that len(first) >= len(second)
	scalarProduct := make([]int64, len(first.Content)-len(second.Content)+1)
	for _, letter := range alphabet {
		// Create 0 and 1 string
		var a, b []*big.Int
		for _, symbol := range first.Content {
			if string(symbol) == letter {
				a = append(a, big.NewInt(1))
			} else {
				a = append(a, big.NewInt(0))
			}
		}
		for _, symbol := range second.Content {
			if string(symbol) == letter {
				b = append(b, big.NewInt(1))
			} else {
				b = append(b, big.NewInt(0))
			}
		}
		reversedB := b
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			reversedB[i], reversedB[j] = reversedB[j], reversedB[i]
		}

		polyA := fft.NewPolynomial(a)
		reversedPolyB := fft.NewPolynomial(reversedB)
		polyA.Mul(polyA, reversedPolyB)
		for i := 0; i < len(scalarProduct); i++ {
			l := len(b)
			j := l - 1 + i
			if j <= len(polyA.StringCoefficients())-4-1 {
				scalarProduct[i] += polyA.Get(j).Int64()
			}
		}
	}
	var maxCoincidence int64
	for _, coincidence := range scalarProduct {
		maxCoincidence = max(maxCoincidence, coincidence)
	}
	if float64(len(first.Content))*precision <= float64(maxCoincidence) {
		fmt.Printf("%s and %s coincide\n", first.Name, second.Name)
	}
}
