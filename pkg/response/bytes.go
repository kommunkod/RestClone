package response

import "net/http"

func ByteResponse(w http.ResponseWriter, statusCode int, contentType string, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write(body)
}
