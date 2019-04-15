package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Handler destroys all of mankind in one clean swoop.
func Handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	limit := 69420
	if h := params.Get("horniness"); h != "" {
		hi, err := strconv.Atoi(h)
		if err != nil {
			fmt.Fprintf(w, "An error occurred: %v", err)
			return
		}
		limit = hi
	}

	for i := 0; i < limit; i++ {
		io.WriteString(w, "UwU ")
	}
}
