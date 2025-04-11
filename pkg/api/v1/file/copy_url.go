package file

import (
	"encoding/json"
	"net/http"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Copy URL
// @Summary Copy URL to destination filesystem
// @Description Copy URL to destination filesystem
// @Tags File
// @Accept json
// @Produce json
// @Param copyURLRequest body rclone.CopyURLRequest true "Copy URL Request"
// @Success 200 {string} string "File copied successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/file/copyUrl [post]
func CopyURL(w http.ResponseWriter, r *http.Request) {
	var request rclone.CopyURLRequest
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

	dst, err := operations.CopyURL(r.Context(), tfs, request.Path, request.URL, request.AutoFilename, request.DstFilenameFromHeader, request.NoClobber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respitem := rclone.CopyUrlData{
		File: rclone.FileItem{
			Name:    dst.Fs().Name(),
			Size:    dst.Size(),
			ModTime: dst.ModTime(r.Context()),
		},
	}

	response.JsonResponse(w, http.StatusOK, respitem)
}
