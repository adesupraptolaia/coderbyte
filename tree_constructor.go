package main
import "fmt"

func TreeConstructor(strArr []string) string {
	monitorParent := map[int][]int{}
	monitorChild := map[int]bool{}

	for _, v := range strArr {
		var d1, d2 int
		fmt.Sscanf(v, "(%d,%d)", &d1, &d2)

		if len(monitorParent[d2]) == 0 {
			monitorParent[d2] = []int{d1}
		} else if len(monitorParent[d2]) == 1 {
			monitorParent[d2] = append(monitorParent[d2], d1)
		} else {
			return "false"
		}

		if monitorChild[d1] {
			return "false"
		}

		monitorChild[d1] = true
	}

	// code goes here
	return "true"
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(TreeConstructor(readline()))

}
