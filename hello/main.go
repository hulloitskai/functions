package main

import (
	"io"
	"net/http"
)

// Handler greets you.
func Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}
