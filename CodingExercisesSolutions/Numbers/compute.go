package main

import "fmt"

func sumTwoNumbers(numbers []int, m map[int]int, sum int) int {
	var answer int

	for _, currentNumber := range numbers {

		value, ok := m[sum-currentNumber]

		if ok {
			answer = value * currentNumber
			ans := fmt.Sprintf("The two Number are %d and %d ", value, currentNumber)
			fmt.Println(ans)
			//break
		}

	}
	return answer
}

func sumThreeNumbers(numbers []int, m map[int]int, sum int) int {
	var answer int

	for index, _ := range numbers {

		yindex := index + 1

		for i, _ := range numbers[yindex:] {

			requiredNumber := sum - (numbers[index] + numbers[i])

			value, ok := m[requiredNumber]

			if ok {
				answer = value * numbers[index] * numbers[i]
				ans := fmt.Sprintf("The three Numbers are %d , %d and %d ", value, numbers[index], numbers[i])
				fmt.Println(ans)
				break
			}

		}

	}
	return answer
}

func changeMap(m map[int]int) {

	m[1234566] = 123456
}
