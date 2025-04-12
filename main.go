package main

import (
	"fmt"
	"math"
)

// Factorial
func getFactorial(n int) int {
	// worst case return 1
	if n < 0 {
		return 1
	}
	factorial := 1
	for i := 1; i < n+1; i++ {
		factorial *= i
	}
	return factorial
}

// check whether an integer is a prime or not
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
	upperLimit := 10
	allPrimesGen := genPrimes(upperLimit)
	fmt.Printf("all primes :- %v\n", allPrimesGen)
	fmt.Printf("factorial for :- %v  ---- res :- %v\n", upperLimit, getFactorial(upperLimit))
}
