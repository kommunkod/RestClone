package main_test

import (
	"testing"
	"time"

	"encoding/json"

	"github.com/clysec/greq"
	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/server"
)

var FtpRemoteConfig = rclone.RemoteConfiguration{
	Name: "local",
	Type: "ftp",
	Parameters: map[string]interface{}{
		"host": "localhost",
		"user": "testuser",
		"pass": "testpassword",
	},
}

func TestFtpList(t *testing.T) {
	t.Log("Starting RestClone Server....")

	err := server.RunBackground(":8910", ":8911")
	if err != nil {
		t.Fatal(err)
	}

	ClearDirectory(t)
	CreateFile(t, "a.txt", "Hello World")

	t.Log("Testing basic functionality...")
	rpr := rclone.RemotePathRequest{
		Remote: FtpRemoteConfig,
	}

	req, err := greq.PostRequest("http://localhost:8910/api/v1/dir/list").WithJSONBody(rpr, nil).Execute()
	if err != nil {
		t.Fatal(err)
	}

	if req.StatusCode != 200 {
		body, err := req.BodyString()
		if err != nil {
			t.Fatal(err)
		}
		t.Fatalf("Expected status 200, got %d. \nResponse Body: %s", req.StatusCode, body)
	}

	var resp rclone.ListFilesResponse
	err = req.BodyUnmarshalJson(&resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Data.Total != 1 {
		t.Fatalf("Expected 1 file, got %d", resp.Data.Total)
	}

	if resp.Data.Files[0].Name != "a.txt" {
		t.Fatalf("Expected file name 'a.txt', got '%s'", resp.Data.Files[0].Name)
	}

	t.Log("Basic functionality test passed.")
}

func TestFtpRead(t *testing.T) {
	t.Log("Starting RestClone Server....")

	err := server.RunBackground(":8910", ":8911")
	if err != nil {
		t.Fatal(err)
	}

	ClearDirectory(t)
	CreateFile(t, "a.txt", "Hello World")

	t.Log("Testing basic functionality...")
	rpr := rclone.ReadFileRequest{
		Remote: FtpRemoteConfig,
		Path:   "a.txt",
	}

	req, err := greq.PostRequest("http://localhost:8910/api/v1/file/read").WithJSONBody(rpr, nil).Execute()
	if err != nil {
		t.Fatal(err)
	}

	if req.StatusCode != 200 {
		body, err := req.BodyString()
		if err != nil {
			t.Fatal(err)
		}
		t.Fatalf("Expected status 200, got %d. \nResponse Body: %s", req.StatusCode, body)
	}

	content, err := req.BodyString()
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "Hello World"
	if content != expectedContent {
		t.Fatalf("Expected file content '%s', got '%s'", expectedContent, content)
	}

	t.Log("Basic functionality test passed.")
}

func TestFtpWrite(t *testing.T) {
	t.Log("Starting RestClone Server....")

	err := server.RunBackground(":8910", ":8911")
	if err != nil {
		t.Fatal(err)
	}

	ClearDirectory(t)

	t.Log("Testing basic functionality...")

	req, err := greq.PostRequest("http://localhost:8910/api/v1/file/write").
		WithMultipartFormBody([]*greq.MultipartField{
			greq.NewMultipartField("file").WithStringValue("Hello World from RestClone").WithFilename("b.txt"),
			greq.NewMultipartField("remote").WithStringValue(func() string {
				data, _ := json.Marshal(FtpRemoteConfig)
				return string(data)
			}()),
			greq.NewMultipartField("path").WithStringValue("/"),
		}).
		Execute()
	if err != nil {
		t.Fatal(err)
	}

	if req.StatusCode != 200 {
		body, err := req.BodyString()
		if err != nil {
			t.Fatal(err)
		}
		t.Fatalf("Expected status 200, got %d. \nResponse Body: %s", req.StatusCode, body)
	}

	time.Sleep(1 * time.Second) // Wait introduced by FTP/SFTP server used for testing

	content := ReadFile(t, "b.txt")
	expectedContent := "Hello World from RestClone"
	if content != expectedContent {
		t.Fatalf("Expected file content '%s', got '%s'", expectedContent, content)
	}

	t.Log("Basic functionality test passed.")
}
