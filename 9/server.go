package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("index.html")
	content, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
}

func readFile(s string) {
	panic("unimplemented")
}

func main() {
	fmt.Println("Listening to port 3000")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
