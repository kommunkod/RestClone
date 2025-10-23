package main

// #cgo CFLAGS: -g3
// #cgo CXXFLAGS: -g3
// #include <stdlib.h>

import "C"

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/kommunkod/restclone/pkg/api/v1"
)

func main() {}

type RequestFormat struct {
	URL     *string
	Method  *string
	Headers *map[string][]string
	Body    string
}

func (req RequestFormat) getRequest() (*http.Request, error) {
	method := "GET"
	if req.Method != nil {
		method = *req.Method
	}

	url := ""
	if req.URL != nil {
		url = *req.URL
	}

	bodyReader := bytes.NewReader([]byte(req.Body))

	httpReq, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	if req.Headers != nil {
		if httpReq.Header == nil {
			httpReq.Header = http.Header{}
		}
		for k, v := range *req.Headers {
			for _, vx := range v {
				httpReq.Header.Add(k, vx)
			}
		}
	}

	return httpReq, nil
}

type ResponseFormat struct {
	StatusCode int
	StatusInfo string
	Headers    map[string][]string
	Body       string
}

func (resp *ResponseFormat) fromHTTPResponse(httpResp *mockResponseWriter) error {
	resp.StatusCode = httpResp.status
	resp.StatusInfo = http.StatusText(httpResp.status)
	resp.Headers = httpResp.header

	resp.Body = base64.StdEncoding.EncodeToString(httpResp.buf.Bytes())

	return nil
}

func (resp *ResponseFormat) Serialize() (string, error) {
	encoded, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encoded), nil
}

type mockResponseWriter struct {
	header http.Header
	buf    bytes.Buffer
	status int
}

func (m *mockResponseWriter) Header() http.Header {
	return m.header
}

func (m *mockResponseWriter) Write(b []byte) (int, error) {
	return m.buf.Write(b)
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.status = statusCode
}

func NewMockWriter() *mockResponseWriter {
	return &mockResponseWriter{
		header: http.Header{},
		buf:    bytes.Buffer{},
		status: http.StatusOK,
	}
}

func Deserialize(s string) RequestFormat {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return RequestFormat{}
	}

	var reqFormat RequestFormat
	err = json.Unmarshal(decoded, &reqFormat)
	if err != nil {
		panic(err)
	}

	return reqFormat
}

//export Operation
func Operation(s string) *C.char {
	fmt.Println(s)
	req := Deserialize(s)
	fmt.Println("Received request for URL:", req)

	httpReq, err := req.getRequest()
	if err != nil {
		resp := ResponseFormat{
			StatusCode: 500,
			StatusInfo: "An error occured: " + err.Error(),
			Headers:    map[string][]string{},
			Body:       "",
		}

		serialized, err := resp.Serialize()
		if err != nil {
			fmt.Println("Error serializing response:", err)
			return C.CString("Error serializing response: " + err.Error())
		}

		return C.CString(serialized)
	}

	fmt.Println("Processing request for URL:", httpReq.URL.String())

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	v1.RegisterRoutes(apiRouter)

	mockWriter := NewMockWriter()

	router.ServeHTTP(mockWriter, httpReq)

	resp := ResponseFormat{}
	resp.fromHTTPResponse(mockWriter)

	serialized, err := resp.Serialize()
	if err != nil {
		fmt.Println("Error serializing response:", err)
		return C.CString("Error serializing response: " + err.Error())
	}

	return C.CString(serialized)
}
