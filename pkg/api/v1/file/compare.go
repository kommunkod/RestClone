package file

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Compare
// @Summary Check if two files are equal
// @Description Check if two files are equal
// @Tags File
// @Accept json
// @Produce json
// @Param checkEqualRequest body rclone.CheckEqualRequest true "Check Equal Request"
// @Success 200 {string} string "Checked successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/file/compare [post]
func Compare(w http.ResponseWriter, r *http.Request) {
	var request rclone.CheckEqualRequest
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

	srcFsFile, err := srcFs.NewObject(r.Context(), request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFsFile, err := dstFs.NewObject(r.Context(), request.DestinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	equal, _, err := operations.CheckHashes(r.Context(), srcFsFile, dstFsFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := rclone.CompareData{
		Equal: equal,
	}

	response.JsonResponse(w, http.StatusOK, resp)
}
