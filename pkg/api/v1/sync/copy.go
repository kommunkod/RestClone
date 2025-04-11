package sync

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/sync"
)

// Sync CopyDir
// @Summary Sync CopyDir
// @Description Sync CopyDir
// @Tags Sync
// @Accept json
// @Produce json
// @Param syncCopyDirRequest body rclone.SyncCopyDirRequest true "Sync CopyDir Request"
// @Success 200 {string} string "Synced successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/sync/copy [post]
func Copy(w http.ResponseWriter, r *http.Request) {
	var request rclone.SyncCopyDirRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystemAtPath(r, w, request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystemAtPath(r, w, request.DestinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = sync.CopyDir(r.Context(), srcFs, dstFs, request.CopyEmptyDirs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ByteResponse(w, http.StatusOK, "text/plain", []byte("Synced successfully"))
}
