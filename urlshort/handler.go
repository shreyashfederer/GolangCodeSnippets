package main

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// URLShort ios
type URLShort struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return nil
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//

var yamlData = `
- path: /wilken
  url: https://wilkenrivera.com/about
- path: /youtube
  url: https://www.youtube.com/watch?v=N9ZWy7xCQ8U&t
- path: /shreyash
  url: https://www.linkedin.com/in/shreyash-bukkawar-816bb1116/
`

//ParseYAML is
func ParseYAML(yamldata []byte) ([]URLShort, error) {

	var pathURL []URLShort

	err := yaml.Unmarshal(yamldata, &pathURL)

	if err != nil {
		return nil, err
	}
	//fmt.Println(pathURL)
	return pathURL, err

}

func buildMap(parsedYaml []URLShort) map[string]interface{} {

	var pathURLs = make(map[string]interface{})

	for _, value := range parsedYaml {

		pathURLs[value.Path] = value.URL

	}

	return pathURLs

}

//RequiredURL is
func RequiredURL(req *http.Request, w http.ResponseWriter) (string, error) {

	parsedYaml, err := ParseYAML([]byte(yamlData))

	if err != nil {

		return "", err
	}

	//fmt.Println(parsedYaml)
	parsedMap := buildMap(parsedYaml)

	if req.URL.Path == "/" {
		fmt.Println("Empty Path in URL")
		// fmt.Fprintf(w, "Empty Path in URL")
		// handler := http.HandlerFunc(fallback)
		// handler.ServeHTTP(w, req)

		fallback(w, req)
		return "", errors.New("Empty Path in URL")

	}

	redirectURL, isURLExists := parsedMap[req.URL.Path].(string)

	// fmt.Println("redirectURL is : ", redirectURL)

	if isURLExists == false {

		//fmt.Fprintf(w, "Short URL doesnt exist in yaml file")
		fallback(w, req)
		return "", errors.New("Short URL doesnt exist in yaml file")
	}
	//fmt.Println(parsedMap[req.URL.Path].(string))
	return redirectURL, err
}

func fallback(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "The path is not present in YAMl file. Please try another path")

}

//YAMLHandler is
func YAMLHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this...
	redirectURL, err := RequiredURL(req, w)

	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println("redirectURL is", redirectURL)
	http.Redirect(w, req, redirectURL, http.StatusFound)
	// redirectURL, isURLExists := parsedMap[req.URL.Path].(string)

	// if isURLExists == false {

	// 	fmt.Fprintf(w, "Short URL doesnt exist in yaml file")

	// 	return
	// }
	// fmt.Println(parsedMap[req.URL.Path].(string))
	//w.WriteHeader(http.StatusOK)

	//http.Redirect(w, req, "www.google.com", http.StatusFound)

	// fmt.Println(req.URL.Path)

	// fmt.Println()

	// fmt.Fprintf(w, req.URL.Path)

	return
}

func main() {

	http.HandleFunc("/", YAMLHandler)
	http.ListenAndServe(":8080", nil)

}
