package sync

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/sync"
)

// Sync MoveDir
// @Summary Sync MoveDir
// @Description Sync MoveDir
// @Tags Sync
// @Accept json
// @Produce json
// @Param syncMoveDirRequest body rclone.SyncMoveDirRequest true "Sync MoveDir Request"
// @Success 200 {string} string "Synced successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/sync/move [post]
func Move(w http.ResponseWriter, r *http.Request) {
	var request rclone.SyncMoveDirRequest
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

	err = sync.MoveDir(r.Context(), srcFs, dstFs, request.DeleteEmptySrcDirs, request.CopyEmptyDirs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ByteResponse(w, http.StatusOK, "text/plain", []byte("Synced successfully"))
}
