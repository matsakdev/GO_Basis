package main

import (
	"fmt"
	"math"
)

func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func findNumbersWithDigitsSum(n, k int) {
	lowerBound := int(math.Pow10(n - 1)) // 10**(n-1), when n==3 lowerBound is 10**2
	upperBound := int(math.Pow10(n)) - 1 // 10**(1)-1, when n==3 upperBound is 999

	for num := lowerBound; num <= upperBound; num++ {
		if sumOfDigits(num) == k {
			fmt.Println(num)
		}
	}
}
