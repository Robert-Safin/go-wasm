package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	t, _ := io.ReadAll(r.Body)
	fmt.Println(string(t))

	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
