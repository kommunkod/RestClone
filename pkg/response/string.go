package response

import "net/http"

func StringResponse(w http.ResponseWriter, statusCode int, body string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte(body))
}
