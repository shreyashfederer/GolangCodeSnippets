package unittesting

import "fmt"

//Hello returns hello world string
func Hello(input string) string {

	output := ""

	if input == "shreyash" {

		output = fmt.Sprintf("%s Bukkawar", input)
	} else if input == "Wilken" {

		output = fmt.Sprintf("%s Rivera", input)

	} else {
		output = input
	}

	return "hello " + output

}

// func main() {

// 	input := "shreyash"

// 	output := Hello(input)

// 	fmt.Println(output)
// }
