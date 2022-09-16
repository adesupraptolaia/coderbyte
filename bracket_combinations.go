package main

import "fmt"

func factorial(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n -= 1
	}

	return result
}

// (n k) = n!/k!(n-k!)
func combination(n, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func BracketCombinations(num int) int {
	return combination(2*num, num) - combination(2*num, num+1)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(BracketCombinations(readline()))

}
