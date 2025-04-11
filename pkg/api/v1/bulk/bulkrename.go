package bulk

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Bulk Rename Files
// @Summary Bulk Rename Files
// @Description Bulk Rename Files
// @Tags Bulk
// @Accept json
// @Produce json
// @Param remote body rclone.BulkRenameFilesRequest true "Remote Configuration"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/bulk/rename [post]
func Rename(w http.ResponseWriter, r *http.Request) {
	var request rclone.BulkRenameFilesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respData := rclone.BulkRenameFilesData{
		RenamedFiles: map[string]string{},
		Errors:       map[string]string{},
	}

	for from, to := range request.NameMap {
		err := operations.MoveFile(r.Context(), tfs, tfs, from, to)
		if err != nil {
			respData.Errors[from] = err.Error()
			continue
		}

		respData.RenamedFiles[from] = to
	}

	response.JsonResponse(w, http.StatusOK, respData)
}
