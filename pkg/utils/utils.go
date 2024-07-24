package utils

import (
    "encoding/json"
	"net/http"
)

// EncodeJSON encodes any data structure into a JSON byte slice
func EncodeJSON(data interface{}) ([]byte, error) {
    return json.Marshal(data)
}

// DecodeJSON decodes a JSON byte slice into a provided data structure
func DecodeJSON(data []byte, target interface{}) error {
    return json.Unmarshal(data, target)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}