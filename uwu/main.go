package main

import (
	"io"
	"net/http"
)

// Handler destroys all of mankind in one clean swoop.
func Handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 69420; i++ {
		io.WriteString(w, "UwU ")
	}
}
