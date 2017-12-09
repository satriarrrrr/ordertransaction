package helpers

import (
	"encoding/json"
	"net/http"
)

// APIResponse response API
type APIResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

// ResponseJSON write response in JSON format
func ResponseJSON(w http.ResponseWriter, message interface{}, code int, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(APIResponse{
		Code:    code,
		Message: message,
	})
}
