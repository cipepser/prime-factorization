package main

import(
	"fmt"
	"math"
)

func divWhilePossible(n, prime int) (newN int, result []int) {	
	var rem int

	for {
		rem = n % prime
		n = n / prime
		
		if rem != 0 {
			break
		} else {
			result = append(result, prime)
		}
	}
	
	newN = n * prime + rem
	return
}

func main() {
	m := 1
	n := 12
	// var result, newResult []int
	var result, newResult []int

	n, newResult = divWhilePossible(n, 2)
	result = append(result, newResult...)

	n, newResult = divWhilePossible(n, 3)
	result = append(result, newResult...)
	
	limit := int(math.Sqrt(float64(n)))
	for n > limit {
		n, newResult = divWhilePossible(n, 6 * m - 1)
		result = append(result, newResult...)

		n, newResult = divWhilePossible(n, 6 * m + 1)
		result = append(result, newResult...)
		
		m++
	} 
	
	fmt.Println(result)
}