package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡first page in go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor escuchando en http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
