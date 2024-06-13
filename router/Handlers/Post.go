package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	data := requestData["data"]

	fmt.Printf("Received data: %s\n", data)

	fmt.Fprintf(w, "Received data: %s", data)
}
