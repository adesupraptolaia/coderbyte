package main

import (
	"fmt"
	"math"
)

func MinWindowSubstring(strArr []string) string {
	n := strArr[0]
	k := strArr[1]

	if k == "" {
		return ""
	}

	countOfK := map[rune]int{}
	for _, v := range k {
		countOfK[v] += 1
	}

	res := []int{-1, -1}
	resLen := math.MaxInt

	have := 0
	need := len(countOfK)
	windowOfN := map[rune]int{}
	left := 0

	for right, v := range n {
		windowOfN[v] += 1
    
		if countOfK[v] > 0 && countOfK[v] == windowOfN[v] {
			have += 1
		}

		for have == need {
			if resLen > (right - left) {
				resLen = right - left + 1
				res = []int{left, right}
			}

			chrLeft := rune(n[left])
			windowOfN[chrLeft] -= 1
			if countOfK[chrLeft] > 0 && windowOfN[chrLeft] < countOfK[chrLeft] {
				have -= 1
			}
			left += 1
		}
	}

	result := n[res[0] : res[1]+1]
  
  // there is no valid substring
	if res[0] == -1 {
		result = ""
	}
  
	return result
}

func main() {
	fmt.Println(MinWindowSubstring(readline()))
}
