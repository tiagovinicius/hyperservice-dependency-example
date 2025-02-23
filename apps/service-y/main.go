package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler para a rota /data
func dataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Service Y")
}

func main() {
	http.HandleFunc("/data", dataHandler)

	fmt.Println("🚀 Server is running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}