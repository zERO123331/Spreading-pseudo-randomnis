package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		threshold := float64(i) / 10
		for j := 0; j < 10; j++ {
			fmt.Println("Seed is: ", j)
			fmt.Println("Threshold is: ", threshold)
			fmt.Println(extendingrandomnisf64(threshold, int64(j)))
		}
	}
}

func extendingrandomnisf64(threashold float64, Seed int64) bool {
	r := rand.New(rand.NewSource(Seed))
	randomness := r.Float64()
	if randomness > threashold {
		return true
	} else {
		return false
	}
}
