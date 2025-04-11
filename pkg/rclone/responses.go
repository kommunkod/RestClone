package rclone

import (
	"time"

	"github.com/rclone/rclone/fs/operations"
)

type ResponseType interface {
	MarshalJSON() ([]byte, error)
	MarshalYAML() ([]byte, error)
}

type Response[T any] struct {
	Success bool   `json:"success" yaml:"success"`
	Error   *Error `json:"error" yaml:"error"`
	Data    T      `json:"data" yaml:"data"`
	// MarshalJSON() ([]byte, error)
	// MarshalYAML() ([]byte, error)
}

type Error struct {
	Code    int    `json:"code" yaml:"code"`
	Message string `json:"message" yaml:"message"`
}

type ListFilesResponse Response[ListFilesData]
type ListFilesData struct {
	Files []operations.ListJSONItem `json:"files"`
	Total int64                     `json:"total"`
}

type ReadFileResponse Response[ReadFileData]
type ReadFileData struct {
	Files []FileItem `json:"files"`
}

type FileItem struct {
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	MimeType string    `json:"mimeType"`
	ModTime  time.Time `json:"modTime"`
}

type WriteFileResponse Response[WriteFileData]
type WriteFileData struct {
	Files []FileItem `json:"files"`
}

type BulkRenameFilesResponse Response[BulkRenameFilesData]
type BulkRenameFilesData struct {
	RenamedFiles map[string]string `json:"renamedFiles"`
	Errors       map[string]string `json:"errors"`
}

type CompareResponse Response[CompareData]
type CompareData struct {
	Equal bool `json:"equal"`
}

type CopyUrlResponse Response[CopyUrlData]
type CopyUrlData struct {
	File FileItem `json:"file"`
}
