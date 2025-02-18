package json

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ResponseWithJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}
}

func ResponseWithError(w http.ResponseWriter, errMessage string, statusCode int) {
	errorResponse := Response{
		Message: errMessage,
		Code:    statusCode,
	}

	ResponseWithJson(w, errorResponse, statusCode)
}

