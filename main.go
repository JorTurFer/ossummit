package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.HandlerFunc(handleRoot))
	http.HandleFunc("/licenses", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/licenses.html")
	})
	http.HandleFunc("/ssl-license", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/ssl-license.html")
	})
	fmt.Println("Listening port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := os.ReadFile("./assets/picture.png")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}
