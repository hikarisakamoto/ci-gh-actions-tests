package main

import "fmt"

func main() {
	fmt.Printf("6 plus 9 is not 69... It is: %d", Sum(6, 9))
}

func Sum(a, b int) int {
	return a + b
}
