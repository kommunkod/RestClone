package file

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Copy File
// @Summary Copy File
// @Description Copy File
// @Tags File
// @Accept json
// @Produce json
// @Param copyFileRequest body rclone.CopyFileRequest true "Copy File Request"
// @Success 200 {string} string "File copied successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/file/copy [post]
func Copy(w http.ResponseWriter, r *http.Request) {
	var request rclone.CopyFileRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.CopyFile(r.Context(), srcFs, dstFs, request.DestinationPath, request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ByteResponse(w, http.StatusOK, "text/plain", []byte("File copied successfully"))
}
