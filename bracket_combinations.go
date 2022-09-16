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

// (2n n) ==> all of posibility, 
// actualy 2n!/n!n! ==>  because you have 2n brackets, and you have n open_bracket and n close_bracket
// ===============
// (2n n+1) ==> all of bad bracket_combinations
// you just need take n+1 bracket (from 2n bracket) to get bad bracket_combination
func BracketCombinations(num int) int {
	return combination(2*num, num) - combination(2*num, num+1)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(BracketCombinations(readline()))

}
