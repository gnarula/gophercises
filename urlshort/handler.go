package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type pair struct {
	Path string
	Url  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}
		redirectHandler := http.RedirectHandler(dest, http.StatusFound)
		redirectHandler.ServeHTTP(w, r)
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
	parsedYAML, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYAML)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]pair, error) {
	var pairs []pair
	err := yaml.Unmarshal(yml, &pairs)
	if err != nil {
		return nil, err
	}
	return pairs, nil
}

func buildMap(pairs []pair) map[string]string {
	dict := make(map[string]string)
	for _, p := range pairs {
		dict[p.Path] = p.Url
	}
	return dict
}
