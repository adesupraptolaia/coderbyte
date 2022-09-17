package main
import (
	"fmt"
	"strings"
  "strconv"
)

func FindIntersection(strArr []string) string {
	a := strings.Split(strArr[0], ", ")
	b := strings.Split(strArr[1], ", ")

	collect := map[string]bool{}

	var shortArr, longArr []string
	if len(a) < len(b) {
		shortArr = a
		longArr = b
	} else {
		shortArr = b
		longArr = a
	}

	maxOfShortArr, _ := strconv.Atoi(shortArr[len(shortArr)-1])

	for _, str := range shortArr {
		collect[str] = true
	}

	result := ""
	for _, strB := range longArr {
		if collect[strB] {
			result += (strB + ",")
		}

		cur, _ := strconv.Atoi(strB)
		if cur > maxOfShortArr {
			break
		}
	}

	if result == "" {
		return "false"
	}

	return result[:len(result)-1] //delete suffix ","
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(FindIntersection(readline()))

}
