package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log.Printf("handler staretd")
	defer log.Printf(("handler ended"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}