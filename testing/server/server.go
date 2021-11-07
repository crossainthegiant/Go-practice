package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/double", doubleHandler)
	log.Fatal(http.ListenAndServe("localhost:4050", nil))
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("v")

	if text == "" {
		http.Error(w, "no words", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "no words", http.StatusBadRequest)
	}

	if _, err := fmt.Fprintln(w, v*2); err != nil {
		http.Error(w, "no response", http.StatusBadRequest)
		return
	}
}
