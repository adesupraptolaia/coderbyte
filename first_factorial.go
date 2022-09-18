package main
import "fmt"

func FirstFactorial(num int) int {
  result := 1
  if num == 0 || num == 1 {
    return result
  }

  for i := 2; i <= num; i++ {
    result *= i
  }
    
  return result;
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(FirstFactorial(readline()))

}
