package file

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Delete File
// @Summary Delete File
// @Description Delete File
// @Tags File
// @Accept json
// @Produce text/plain
// @Param remote body rclone.DeleteFileRequest true "Remote Configuration"
// @Success 200 {string} string "File deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/file/delete [post]
func Delete(w http.ResponseWriter, r *http.Request) {
	var request rclone.DeleteFileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tfo, err := tfs.NewObject(r.Context(), request.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.DeleteFile(r.Context(), tfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ByteResponse(w, http.StatusOK, "text/plain", []byte("File deleted successfully"))
}
