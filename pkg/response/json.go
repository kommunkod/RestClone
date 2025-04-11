package response

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
)

func JsonResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func JsonSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := rclone.Response[any]{
		Success: true,
		Error:   nil,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
