package main

import (
	"flag"
	"fmt"
	"goexcrs/urlshort"
	"io/ioutil"
	"net/http"
)

var (
	yamlFile = flag.String("yaml", "", "Accepts the filename of a YAML file from wich loading the URLs")
)

func main() {
	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	if *yamlFile != "" {

		yamlData, err := readYAMLFile(*yamlFile)
		if err != nil {
			yamlData = populateYAML()
		}
		yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
		if err != nil {
			panic(err)
		}
		serveHandler(yamlHandler)
	} else {
		jsonData := []byte(`{"/yt":"https://youtube.com", "/me":"https://medium.com"}`)
		jsonHandler, err := urlshort.JSONHandler(jsonData, mapHandler)
		if err != nil {
			panic(err)
		}
		serveHandler(jsonHandler)
	}

}

func serveHandler(h http.Handler) {
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", h)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readYAMLFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func populateYAML() []byte {
	return []byte(`
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	  url: https://github.com/gophercises/urlshort/tree/solution
	`)
}
