package unittesting

import (
	"fmt"
	"testing"
)

func init() {

	fmt.Println("This line would excute before the program starts")
}

func TestHello(t *testing.T) {

	type test struct {
		name     string
		input    string
		expected string
	}

	tests := []test{

		{

			name: "Print Hello World",

			input: "world",

			expected: "hello world",
		},
		{

			name: "Print Hello shreyash",

			input: "shreyash",

			expected: "hello shreyash Bukkawar",
		},

		{

			name: "Print Hello Wilken",

			input: "Wilken",

			expected: "hello Wilken Rivera",
		},
	}

	for index, testCase := range tests {

		//SetUp
		fmt.Printf("Setup after each test case %d\n", index)

		t.Run(testCase.name, func(t *testing.T) {

			actualOutput := Hello(testCase.input)

			if actualOutput != testCase.expected {

				t.Errorf("expected %s , got %s", testCase.expected, actualOutput)
			}

		})

	}

}
