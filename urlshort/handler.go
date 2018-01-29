package urlshort

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-yaml/yaml"
)

type ShortURL struct {
	ShortPath string `yaml:"path"`
	RealPath  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
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
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	shortURLs, err := ParseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathsToUrls := sliceToMapShortURLs(shortURLs)
	return MapHandler(pathsToUrls, fallback), nil
}

func ParseYAML(data []byte) ([]ShortURL, error) {
	var shortURLs []ShortURL
	err := yaml.Unmarshal([]byte(data), &shortURLs)
	if err != nil {
		log.Fatalf("cannot unmarshal YAML data: %v", err)
		return nil, err
	}
	return shortURLs, nil
}

func sliceToMapShortURLs(shortURLs []ShortURL) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, u := range shortURLs {
		pathsToUrls[u.ShortPath] = u.RealPath
	}
	return pathsToUrls
}

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	shortURLs, err := ParseJSON(jsonData)
	if err != nil {
		return nil, err
	}
	// pathsToUrls := sliceToMapShortURLs(shortURLs)
	return MapHandler(shortURLs, fallback), nil
}

func ParseJSON(data []byte) (map[string]string, error) {
	var shortURLs map[string]string
	err := json.Unmarshal([]byte(data), &shortURLs)
	if err != nil {
		log.Fatalf("cannot unmarshal JSON data: %v", err)
		return nil, err
	}
	fmt.Println(shortURLs)
	return shortURLs, nil
}
