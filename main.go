package main

import (
	"log"
	"math"
)

func main() {
	s := "agadga"

	log.Print(s[:int32(math.Min(float64(len(s)), 20))])
}
