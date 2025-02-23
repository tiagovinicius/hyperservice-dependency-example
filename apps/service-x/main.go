package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	serviceYURL = os.Getenv("SERVICE_Y_API_URL")
)

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func callServiceY(w http.ResponseWriter, r *http.Request) {
	if serviceYURL == "" {
		http.Error(w, `{"message": "❌ SERVICE_Y_API_URL is not set!"}`, http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(serviceYURL)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "❌ Failed to reach SERVICE_Y_API_URL", "error": "%v"}`, err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lendo o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, `{"message": "❌ Failed to read response body"}`, http.StatusInternalServerError)
		return
	}

	// Criando resposta JSON
	response := Response{
		Message: "Data fetched from Service Y",
		Data:    string(body),
	}

	// Configurando resposta como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(response)
}


func main() {
	if serviceYURL == "" {
		log.Fatal("❌ SERVICE_Y_API_URL is not set!")
	}

	http.HandleFunc("/call-service-y", callServiceY)

	fmt.Println("🚀 Server is running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}