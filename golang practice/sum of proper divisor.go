package main

import (
	"fmt"
)

func PrimeFactors(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {

		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func Power(p, i int) int {
	result := 1
	for j := 0; j < i; j++ {
		result *= p
	}
	return result
}

func SumOfProperDivisors(n int) int {
	pfs := PrimeFactors(n)

	m := make(map[int]int)
	for _, prime := range pfs {
		_, ok := m[prime]
		if ok {
			m[prime] += 1
		} else {
			m[prime] = 1
		}
	}

	sumOfAllFactors := 1
	for prime, exponents := range m {
		sumOfAllFactors *= (Power(prime, exponents+1) - 1) / (prime - 1)
	}
	return sumOfAllFactors - n
}

func main() {
	fmt.Println(SumOfProperDivisors(220))
}
