package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFile(m map[int]int) []int {

	file, err := os.Open("file.txt")

	var numbers []int

	if err != nil {

		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		number, err := strconv.Atoi(scanner.Text())
		m[number] = number
		if err != nil {

			log.Fatal(err)

		}
		numbers = append(numbers, number)
	}

	return numbers

}
