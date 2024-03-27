package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 		w.Header().Add("Content-Type", "text/html") // add content type header
		fmt.Fprintf(w, "Hello world!", w)
 	})
 	http.ListenAndServe(":3333", nil)
}
