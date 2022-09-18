package main

import (
	"fmt"
	"strconv"
)

func QuestionsMarks(str string) string {

	result := "false"

	start := 0
	totalQuestionMark := 0

	for i := 0; i < len(str); i++ {
		chr := str[i : i+1]

    // ascii number for 0-9 between 48-57
		if str[i] >= 48 && str[i] <= 57 {
			curNumber, _ := strconv.Atoi(chr)

			if start+curNumber == 10 {
				if totalQuestionMark != 3 {
					return "false"
				} else {
					result = "true"
				}
			}

			totalQuestionMark = 0
			start = curNumber
		} else if chr == "?" {
			totalQuestionMark += 1
		}
	}

	// code goes here
	return result
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(QuestionsMarks(readline()))

}
