package main

import (
	"math"
)

func kem(n int) (int, int) {
	count := 0
	larger := 0.0
	for l := float64(n / 2); l >= 2; l-- {
		for i := float64(2); i < float64(n); i++ {
			x := math.Pow(l, i)
			if x >= float64(n) {
				break
			}
			// fmt.Printf("%.0f**%.0f = %.2f", l, i, math.Pow(l, i))
			if larger < x {
				count = 1
				larger = x
				// fmt.Printf(" - counting 1")
			} else if larger == x {
				count++
				larger = x
				// fmt.Printf(" - counting +")
			}
			// fmt.Println("")
		}
	}

	return int(larger), count
}
