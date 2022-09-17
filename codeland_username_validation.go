package main

import (
	"fmt"
	"unicode"
)


func CodelandUsernameValidation(str string) string {
	l := len(str)

	if l < 4 || l > 25 {
		return "false"
	} else if str[l-1:l] == "_" {
		return "false"
	} else if !unicode.IsLetter(rune(str[0])) {
		return "false"
	}

	for i := 1; i < len(str); i++ {
		chr := rune(str[i])
		if !unicode.IsLetter(chr) && !unicode.IsDigit(chr) && string(str[i]) != "_" {
			return "false"
		}
	}

	return "true"
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(CodelandUsernameValidation(readline()))

}
