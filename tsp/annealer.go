package tsp

import (
	"fmt"
	//"math"
	"math/rand"
)

type Annealer struct{
	Random int
}

func Rand_num() {
	rand.Seed(42)
	fmt.Println("My favorite number is", rand.Intn(10))
}