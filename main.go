package main

import (
	"net/http"
	"imgscale"
)

func main() {
	http.HandleFunc("/", imgscale.Handler)
	http.ListenAndServe(":8080", nil)
}

