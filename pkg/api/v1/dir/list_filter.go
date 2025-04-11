package dir

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/kommunkod/restclone/pkg/rclone"
	"github.com/kommunkod/restclone/pkg/response"
	"github.com/rclone/rclone/fs/operations"
)

// Filtered List Files
// @Summary Filtered List Files
// @Description List files in a given directory with a filter
// @Tags Directory
// @Accept json
// @Produce json
// @Param remote body rclone.FilteredListFilesRequest true "Remote Configuration"
// @Success 200 {object} rclone.ListFilesResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/dir/filterlist [post]
func FilteredList(w http.ResponseWriter, r *http.Request) {
	var request rclone.FilteredListFilesRequest
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

	filterFunc := func(item *operations.ListJSONItem) bool {
		return true
	}

	switch request.FilterType {
	case rclone.FilterTypePrefix:
		filterFunc = func(item *operations.ListJSONItem) bool {
			return strings.HasPrefix(item.Name, request.Filter)
		}
	case rclone.FilterTypeSuffix:
		filterFunc = func(item *operations.ListJSONItem) bool {
			return strings.HasSuffix(item.Name, request.Filter)
		}
	case rclone.FilterTypeRegex:
		filterFunc = func(item *operations.ListJSONItem) bool {
			return regexp.MustCompile(request.Filter).MatchString(item.Name)
		}
	case rclone.FilterTypeWildcard:
		filterFunc = func(item *operations.ListJSONItem) bool {
			return strings.Contains(item.Name, request.Filter)
		}
	}

	err = operations.ListJSON(context.Background(), tfs, request.Path, &request.Options, func(item *operations.ListJSONItem) error {
		if filterFunc(item) {
			outData.Files = append(outData.Files, *item)
		}

		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outData.Total = int64(len(outData.Files))
	response.JsonResponse(w, http.StatusOK, outData)
}
