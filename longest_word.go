package main
import "fmt"

func isNumber(chr byte) bool {
	if chr >= 48 && chr <= 57 {
		return true
	}
	return false
}

func isAlphabeth(chr byte) bool {
	if chr >= 65 && chr <= 90 {
		return true
	}

	if chr >= 97 && chr <= 122 {
		return true
	}

	return false
}

func LongestWord(sen string) string {
	var result, temp string

	for i := 0; i < len(sen); i++ {
		chr := sen[i]

		if isAlphabeth(chr) || isNumber(chr) {
			temp += string(chr)
		} else {
			if len(temp) > len(result) {
				result = temp
			}

			temp = ""
		}

		if len(result) >= (len(sen) - i + len(temp)) {
			return result
		}
	}

	if len(temp) > len(result) {
		result = temp
	}

	return result
}


func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(LongestWord(readline()))

}
