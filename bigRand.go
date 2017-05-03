package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	one := big.NewInt(1)

	for i := 0; i < 20; i++ {
		a := new(big.Int).Rand(rnd, big.NewInt(10))
		a.Add(a, one)
		fmt.Println(a)
	}
}
