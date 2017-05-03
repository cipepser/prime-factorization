package main

import(
	"fmt"
	"math/rand"
	"time"
	"math/big"
)

func isPrime(p *big.Int) bool {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	if p.Cmp(two) == 0 {
		return true
	}
	
	// p - 1 = 2^s * dに分解する
	d := new(big.Int).Sub(p, one)
	s := 0
	for new(big.Int).And(d, one).Cmp(zero) == 0 {
		d.Rsh(d, 1)
		s++
	}
	
	n := new(big.Int).Sub(p, one)
	k := 20
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < k; i++ {
		result := false
		// [1, n-1]からaをランダムに選ぶ
		a := new(big.Int).Rand(rnd, n)
		a.Add(a, one)
		
		// a^{2^r * d} mod p != -1(= p - 1 = n)の比較
		tmp := new(big.Int).Exp(a, d, p)
		for r := 0; r < s; r++ {
			if tmp.Cmp(n) == 0 {
				result = true
				break
			}
			tmp.Exp(tmp, two, p)
		}
		
		// a^d != 1 mod p の比較
		if !result && new(big.Int).Exp(a, d, p).Cmp(one) != 0 {	
			return false
		}
	}
	
	return true
}

func main() {
	fmt.Println("------------Prime-------------")
	primes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	for _, p := range primes {
		fmt.Println(p, ": ", isPrime(big.NewInt(p)))
	}
	
	fmt.Println("------------Composite-------------")
	
	composites := []int64{9, 15, 21, 25, 27, 33, 35, 39, 45, 49, 51, 55, 57, 63, 65, 69, 75, 77, 81, 85, 87, 91, 93, 95, 99}
	for _, c := range composites {
		fmt.Println(c, ": ", isPrime(big.NewInt(c)))
	}
	
	n, _ := new(big.Int).SetString("16329805957987392833", 10) // 7508269669 * 2174909357の合成数
	fmt.Println(n, ": ", isPrime(n))

	n, _ = new(big.Int).SetString("75367631466941141791247084978233913478587466161271622771710895906453753967164967411257954571043148038883968051384064219878882200013408641715083492640764339527637873502061727098279950196988248939994544197535642414553039637030217247486372549770605212095298947963579455290328750806815560590656493190626063424483777638146122548073605596393468295558919015322872295712049815793564477203891171966598995001945926910413702372501638140065699147590309176011776418089018393636130620122226631715476290375904684033", 10) // 500桁の素数
	fmt.Println(n, ": ", isPrime(n))
		
}