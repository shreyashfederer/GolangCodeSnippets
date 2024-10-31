package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	output := make(chan int)

	for index := 1; index <= 5; index += 1 {
		wg.Add(1)
		go squareNumber(index, output, &wg)

	}

	go func() {
		wg.Wait()
		close(output)
	}()

	printOutput(output)

}

func squareNumber(number int, output chan int, wg *sync.WaitGroup) {

	output <- number * number

	wg.Done()

}

func printOutput(output chan int) {

	for value := range output {

		fmt.Println(value)
	}
}
