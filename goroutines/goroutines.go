package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//URLs contains list of url
var URLs = []string{
	"https://www.youtube.com",
	"https://www.google.com",
	"https://www.golang.org",
}

// MultipleRequests calls multiple get Requests using goroutines
func MultipleRequests(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	for index, url := range URLs {

		wg.Add(1)
		go func(url string) {

			fmt.Printf("Inside goRoutine %d \n", index)
			response, err := http.Get(url)

			if err != nil {

				fmt.Fprintf(w, "%+v\n", err)
			}

			fmt.Fprintf(w, "%+v\n", response.Status)
			fmt.Fprintf(w, "%+v\n", response.Request.Method)

			wg.Done()

		}(url)

		fmt.Printf("Last line of iteration number %d \n", index)

	}
	wg.Wait()
}

func main() {

	fmt.Println("Program is started")

	http.HandleFunc("/", MultipleRequests)

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("Last Line of program")

}
