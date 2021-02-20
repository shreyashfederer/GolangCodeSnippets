package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var x interface{} = 7

	i := x.(int)

	fmt.Println(i)

	var answer int

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		arr := strings.Split(scanner.Text(), ":")

		passwordPolicy := arr[0]
		currentpassword := arr[1][1:]

		policy := passwordPolicy[:len(passwordPolicy)-2]

		lowerlimit, _ := strconv.Atoi(strings.Split(policy, "-")[0])
		upperlimit, _ := strconv.Atoi(strings.Split(policy, "-")[1])

		number := string(arr[0][len(arr[0])-1])

		// fmt.Println(policy, number, currentpassword)
		// fmt.Println(lowerlimit, upperlimit)

		count := strings.Count(currentpassword, number)
		//fmt.Println(count)
		if count >= lowerlimit && count <= upperlimit {
			answer = answer + 1
		}

	}

	fmt.Println(answer)
}
