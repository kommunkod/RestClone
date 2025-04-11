package dir

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// List Files
// @Summary List Files
// @Description List Files
// @Tags Directory
// @Accept json
// @Produce json
// @Param remote body rclone.ListFilesRequest true "Remote Configuration"
// @Success 200 {object} rclone.ListFilesResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/dir/list [post]
func List(w http.ResponseWriter, r *http.Request) {
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
	outData := rclone.ListFilesData{}

	err = operations.ListJSON(context.Background(), tfs, request.Path, &request.Options, func(item *operations.ListJSONItem) error {
		outData.Files = append(outData.Files, *item)
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outData.Total = int64(len(outData.Files))
	response.JsonResponse(w, http.StatusOK, outData)
}
