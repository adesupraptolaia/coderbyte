package main
import "fmt"

func FirstReverse(str string) string {
  var result string

  for i := len(str)-1; i >=0; i-- {
    result += str[i:i+1]
  }

  return result;
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(FirstReverse(readline()))

}
