package response

import (
	"io"
	"net/http"
)

func StreamResponse(w http.ResponseWriter, statusCode int, headers map[string][]string, body io.Reader) {
	if headers["Content-Type"] == nil {
		headers["Content-Type"] = []string{"text/plain"}
	}

	for k, v := range headers {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	w.WriteHeader(statusCode)
	io.Copy(w, body)
}
