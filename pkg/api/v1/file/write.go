package file

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/object"
	"github.com/rclone/rclone/fs/operations"
)

// Write File
// @Summary Write File
// @Description Write File. You can attach an arbitrary number of files to the request. All have to be placed in the "file" field.
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Param remote formData string true "Remote Configuration"
// @Param path formData string true "Path"
// @Param overwrite formData string false "Overwrite"
// @Success 200 {object} rclone.WriteFileResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/file/write [put]
func Write(w http.ResponseWriter, r *http.Request) {
	request := rclone.WriteFileRequest{}

	remoteData := r.FormValue("remote")
	if remoteData == "" {
		http.Error(w, "Remote configuration is required", http.StatusBadRequest)
		return
	}

	overwrite := strings.Trim(r.FormValue("overwrite"), "\r\n")
	if slices.Contains([]string{"true", "1", "yes", "y"}, overwrite) {
		request.Overwrite = true
	}

	json.Unmarshal([]byte(remoteData), &request.Remote)

	pathData := strings.Trim(r.FormValue("path"), "\r\n")
	if !strings.HasSuffix(pathData, "/") && pathData != "" && pathData != "/" {
		pathData = pathData + "/"
	}

	request.Path = pathData

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.MultipartForm.File["file"] == nil {
		http.Error(w, "File field 'file' is required", http.StatusBadRequest)
		return
	}

	resp := rclone.WriteFileData{}

	for _, file := range r.MultipartForm.File["file"] {
		fPath := request.Path + file.Filename

		statx, err := operations.StatJSON(r.Context(), tfs, fPath, &operations.ListJSONOpt{})
		if err == nil && statx != nil && !request.Overwrite {
			http.Error(w, "File already exists", http.StatusInternalServerError)
			return
		}

		fo, err := file.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer fo.Close()

		memobj := object.NewMemoryObject(fPath, time.Now(), nil)

		ob, err := tfs.Put(context.Background(), fo, memobj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Files = append(resp.Files, rclone.FileItem{
			Name:     fPath,
			Size:     ob.Size(),
			MimeType: file.Header.Get("Content-Type"),
			ModTime:  ob.ModTime(r.Context()),
		})
	}

	response.JsonResponse(w, http.StatusOK, resp)
}
