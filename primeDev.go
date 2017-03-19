package main

import (
	"fmt"
	"math/big"
)

const (
	DECIMAL int = 10
)

func main()  {
	n, _ := new(big.Int).SetString("16329805957987392833", DECIMAL)

	// 3から2ずつ増やしていく
	q := big.NewInt(3)
	a := big.NewInt(2)
	r := big.NewInt(0)

	zero := big.NewInt(0)
	
	for {
		tmp, _ := new(big.Int).SetString(n.String(), DECIMAL)
		p, r := tmp.QuoRem(tmp, q, r)

		if r.Cmp(zero) == 0 {
			fmt.Printf("p: %v\n", p)
			fmt.Printf("q: %v\n", q)
			fmt.Printf("n: %v\n", n)
			break
		}
		
		q.Add(q, a)
	}
}