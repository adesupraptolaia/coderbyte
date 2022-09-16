package main

import "fmt"

func multiply(n, m int) int {
	result := 1
	for n >= m {
		result *= n
		n -= 1
	}

	return result
}

// (n k) = n!/k!(n-k!)
// = n*(n-1)*...*(max+1) / min*(min-1)*...*1
func combination(n, k int) int {
	max := n - k
	min := k
	if k > n-k {
		max = k
		min = n - k
	}

	return multiply(n, max+1) / multiply(min, 1)
}

// (2n n) ==> all of posibility,
// actualy 2n!/n!n! ==>  because you have 2n brackets, and you have n open_bracket and n close_bracket
// ===============
// (2n n+1) ==> all of bad bracket_combinations
// you just need take n+1 bracket (from 2n bracket) to get bad bracket_combination
func BracketCombinations(num int) int {
	if num == 0 {
		return 1
	}
	return combination(2*num, num) - combination(2*num, num+1)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(BracketCombinations(readline()))

}
