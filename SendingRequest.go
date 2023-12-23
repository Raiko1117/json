package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}
type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST request are allowed", http.StatusMethodNotAllowed)
			return
		}
		var req JsonRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		if req.Message == "" {
			http.Error(w, "Missing or empty 'message' field in JSON", http.StatusBadRequest)
			return
		}
		fmt.Printf("Received message: %s\n", req.Message)

		res := JsonResponse{
			Status:  "success",
			Message: "Data successfully accepted!",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello world")
	})

	fmt.Println("Server is running on:8080")
	http.ListenAndServe(":8080", nil)
}
