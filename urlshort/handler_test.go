package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type parseYAMLtests struct {
	name  string
	input []byte

	expected []URLShort

	wantErr     bool
	expectedErr error
}

type yamlHandlertests struct {
	name string

	input map[string]interface{}

	inputPath string

	expected expectedResponse

	wantErr     bool
	expectedErr error
}

type expectedResponse struct {
	Status int
	URL    string
}

func TestParseYAML(t *testing.T) {

	tests := []parseYAMLtests{

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

func TestYAMLHandler(t *testing.T) {

	tests := []yamlHandlertests{

		{
			name: "check for a Successful Redirect",

			input: map[string]interface{}{

				"/shreyash": "https://www.linkedin.com/in/shreyash-bukkawar-816bb1116",
				"/wilken":   "https://wilkenrivera.com/about",
				"/youtube":  "https://www.youtube.com/watch?v=N9ZWy7xCQ8U&t",
			},
			inputPath: "/wilken",

			expected: expectedResponse{

				Status: 302,
				URL:    "https://wilkenrivera.com/about",
			},

			wantErr: false,
		},

		{
			name: "Give a path that doesnt existin map and get error",

			input: map[string]interface{}{

				"/shreyash": "https://www.linkedin.com/in/shreyash-bukkawar-816bb1116",
				"/wilken":   "https://wilkenrivera.com/about",
				"/youtube":  "https://www.youtube.com/watch?v=N9ZWy7xCQ8U&t",
			},
			inputPath: "/unknownpath",

			expected: expectedResponse{

				Status: 404,
				URL:    "",
			},

			wantErr: false,
		},
	}

	for _, currentTest := range tests {

		t.Run(currentTest.name, func(t *testing.T) {

			//url := currentTest.input[currentTest.inputPath].(string)
			//fmt.Printf("url is %v \n", url)
			req, err := http.NewRequest("GET", currentTest.inputPath, nil)

			if err != nil {
				t.Errorf("Error in request")
			}

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(YAMLHandler)

			//http.DefaultServeMux.ServeHTTP(responseRecorder, req)
			handler.ServeHTTP(responseRecorder, req)
			// fmt.Println("hey: ", responseRecorder)

			// fmt.Printf("response code %d url %v \n ", responseRecorder.Code, responseRecorder.HeaderMap.Get("Location"))

			actualResponse := expectedResponse{

				Status: responseRecorder.Code,

				URL: responseRecorder.HeaderMap.Get("Location"),
			}

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

			if !reflect.DeepEqual(actualResponse, currentTest.expected) {

				t.Errorf("Expected = %v got = %v ", currentTest.expected, actualResponse)
			}

		})

	}

}
