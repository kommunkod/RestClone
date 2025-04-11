package dir

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Remove Dirs (force)
// @Summary Remove Directories
// @Description Recursively remove directories
// @Tags Directory
// @Accept json
// @Produce json
// @Param rmdirsRequest body rclone.RmdirsRequest true "Remove Directories Request"
// @Success 200 {string} string "Removed successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/dir/removeRecursive [post]
func Rmdirs(w http.ResponseWriter, r *http.Request) {
	var request rclone.RmdirsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.Rmdirs(r.Context(), tfs, request.Path, request.LeaveRoot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ByteResponse(w, http.StatusOK, "text/plain", []byte("Removed successfully"))
}
