package file

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Read File
// @Summary Read File
// @Description Read File
// @Tags File
// @Accept json
// @Produce json
// @Param remote body rclone.ReadFileRequest true "Remote Configuration"
// @Success 200 {object} rclone.ReadFileResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/file/read [post]
func Read(w http.ResponseWriter, r *http.Request) {
	var request rclone.ListFilesRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request.Remote.Parameters["fast-list"] = true

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

	if tfo.Size() == 0 || tfo.Size() == -1 {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	stat, err := operations.StatJSON(r.Context(), tfs, request.Path, &operations.ListJSONOpt{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content, err := operations.ReadFile(r.Context(), tfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ByteResponse(w, http.StatusOK, stat.MimeType, content)
}
