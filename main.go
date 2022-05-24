package fork_test

import "fmt"

func main() {
	fmt.Sprintf("fork-test")
	fmt.Sprintf("code change")
	fmt.Sprintf("Another code change")
}

func test1() {
	fmt.Sprintf("fork-test1")
}

func test() {
 	fmt.Sprintf("fork-test")
}
