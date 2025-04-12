package main

import (
	"fmt"
	"math"
)

func isPrime(allPrimes []int, n int) bool {
	nSqrt := int(math.Sqrt(float64(n)))
	for i := 0; len(allPrimes) != 0 && allPrimes[i] <= nSqrt; i += 1 {
		if n%allPrimes[i] == 0 {
			return false
		}
	}
	return true
}

func genPrimes(n int) []int {
	allPrimes := []int{}
	for i := 2; i <= n; i += 1 {
		if isPrime(allPrimes, i) {
			allPrimes = append(allPrimes, i)
		}
	}
	return allPrimes
}

func main() {
	upperLimit := 100
	allPrimesGen := genPrimes(upperLimit)
	fmt.Printf("all primes :- %v\n", allPrimesGen)
}
