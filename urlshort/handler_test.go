package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestParseYAML(t *testing.T) {

	type test struct {
		name  string
		input []byte

		expected []URLShort

		wantErr     bool
		expectedErr error
	}

	tests := []test{

		{

			name:  "Check for valid unmarshalling of data",
			input: []byte(yamlData),
			expected: []URLShort{

				{
					Path: "/wilken",
					URL:  "https://wilkenrivera.com/about",
				},
				{
					Path: "/youtube",
					URL:  "https://www.youtube.com/watch?v=N9ZWy7xCQ8U&t",
				},

				{
					Path: "/shreyash",
					URL:  "https://www.linkedin.com/in/shreyash-bukkawar-816bb1116/",
				},
			},

			wantErr: false,
		},

		{
			name: "Send unformatted Yaml and get error",

			input: []byte(`
			
			- path: /wilken
 			 url: https://wilkenrivera.com/about
			-path: /youtube
  			url: https://www.youtube.com/watch?v=N9ZWy7xCQ8U&t
         	- path: /shreyash
  			  url: https://www.linkedin.com/in/shreyash-bukkawar-816bb1116/
			`),

			expected: []URLShort{},

			wantErr:     true,
			expectedErr: errors.New("found character that cannot start any token"),
		},
	}

	for _, currentTest := range tests {

		t.Run(currentTest.name, func(t *testing.T) {

			actualOutput, err := ParseYAML(currentTest.input)

			if err != nil {
				if currentTest.wantErr {

					if strings.Contains(err.Error(), currentTest.expectedErr.Error()) {
						//Tests Passed if it conatins error message
						return
					}

					t.Errorf("Expected error = %v got = %v ", currentTest.expectedErr.Error(), err.Error())
				}
				return
			}

			if !reflect.DeepEqual(actualOutput, currentTest.expected) {

				t.Errorf("Expected = %v got = %v ", currentTest.expected, actualOutput)
			}

		})

	}

}
