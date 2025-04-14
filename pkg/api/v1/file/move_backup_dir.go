package file

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Move Backup Dir
// @Summary Move Backup Dir
// @Description Move a file to a backup directory
// @Tags File
// @Accept json
// @Produce json
// @Param moveBackupDirRequest body rclone.MoveBackupDirRequest true "Move Backup Dir Request"
// @Success 200 {string} string "Backup directory moved successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/moveBackupDir [post]
func MoveBackupDir(w http.ResponseWriter, r *http.Request) {
	var request rclone.MoveBackupDirRequest
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

	dstFsDir, err := dstFs.NewObject(r.Context(), request.DestinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.MoveBackupDir(r.Context(), srcFs, dstFsDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ByteResponse(w, http.StatusOK, "text/plain", []byte("Backup directory moved successfully"))

}
