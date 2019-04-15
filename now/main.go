package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler responds with the current time.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The time is currently: %s", time.Now())
}
