package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

var serviceYURL string // será atribuída no main()

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, `{"message": "❌ Failed to read response body"}`, http.StatusInternalServerError)
		return
	}

	response := Response{
		Message: "✅ Data fetched from Service Y",
		Data:    string(body),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Println("File .env not found")
	}

	// Assign value after loading .env
	serviceYURL = os.Getenv("SERVICE_Y_API_URL")
	if serviceYURL == "" {
		log.Fatal("❌ SERVICE_Y_API_URL is not set!")
	}

	http.HandleFunc("/call-service-y", callServiceY)

	fmt.Println("🚀 Server is running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}