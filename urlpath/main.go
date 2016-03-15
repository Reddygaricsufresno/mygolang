package main

import (
	"fmt"
	"net/http"
	"log"
)


func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		res.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(res, "name: %v", req.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}