package main

import (
	"fmt"
)

func main() {

	sum := 2020
	//numbers := make([]int, 0)

	m := make(map[int]int)

	numbers := readFile(m)
	answer := sumTwoNumbers(numbers, m, sum)
	product := sumThreeNumbers(numbers, m, sum)

	fmt.Println(m)
	changeMap(m)

	fmt.Println("Product of two Numbers is : ", answer)

	fmt.Println("Product of three Numbers is : ", product)

	fmt.Println(m)

}
