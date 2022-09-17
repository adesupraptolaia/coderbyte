package main
import "fmt"

func BracketMatcher(str string) string {

	openBracket := 0

	for _, v := range str {
		chr := string(v)

		if chr != "(" && chr != ")" {
			continue
		}

		if chr == "(" {
			openBracket += 1
			continue
		}

		openBracket -= 1

		if openBracket < 0 {
			return "0"
		}
	}

	if openBracket > 0 {
		return "0"
	}

	return "1"
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(BracketMatcher(readline()))

}
