package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Message struct {
	Text string `json:"text"`
}

type FactorialResponse struct {
	Input  int `json:"input"`
	Result int `json:"result"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func FactorialHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputStr := params["number"]
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := factorial(input)
	response := FactorialResponse{
		Input:  input,
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Create a new router instance
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/ping", PingHandler).Methods("GET")
	router.HandleFunc("/factorial/{number}", FactorialHandler).Methods("GET")

	// Start the HTTP server
	log.Println("Server started on port 18080")
	log.Fatal(http.ListenAndServe(":18080", router))
}
