package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dataset := map[string]interface{}{
		"id":   1,
		"name": "Arun",
		"age":  30,
		"skills": []string{
			"Go",
			"Python",
			"JavaScript",
		},
	}

	jsonData, err := json.Marshal(dataset)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
