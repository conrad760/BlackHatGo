package first

import "fmt"

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go TDD")
	result := Calculate(2)
	fmt.Println(result)
}