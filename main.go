package fork_test

import "fmt"

func main() {
	fmt.Sprintf("fork-test")
	fmt.Sprintf("code change")
	fmt.Sprintf("Another code change1")
	fmt.Sprintf("upstream change")
}

func test() {
 	fmt.Sprintf("fork-test123")
}

func sam() {
	fmt.Sprintf("sam")
}

func test1() {
	fmt.Sprintf("fork-test1")
}
