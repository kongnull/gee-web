package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
		if err != nil {
			return
		}
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			_, err := fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
			if err != nil {
				return
			}
		}
	})

	err := r.Run(":9200")
	if err != nil {
		return 
	}
}
