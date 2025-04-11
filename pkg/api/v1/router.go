package v1

import (
	"github.com/gorilla/mux"
	"github.com/kommunkod/restclone/pkg/api/v1/bulk"
	"github.com/kommunkod/restclone/pkg/api/v1/dir"
	"github.com/kommunkod/restclone/pkg/api/v1/file"
	"github.com/kommunkod/restclone/pkg/api/v1/sync"
	_ "github.com/rclone/rclone/backend/all"
)

func RegisterRoutes(router *mux.Router) {
	RegisterBulkRoutes(router.PathPrefix("/bulk").Subrouter())
	RegisterDirectoryRoutes(router.PathPrefix("/dir").Subrouter())
	RegisterFileRoutes(router.PathPrefix("/file").Subrouter())
	RegisterSyncRoutes(router.PathPrefix("/sync").Subrouter())
}

func RegisterBulkRoutes(router *mux.Router) {
	router.HandleFunc("/rename", bulk.Rename).Methods("POST")
}

func RegisterDirectoryRoutes(router *mux.Router) {
	router.HandleFunc("/listFilter", dir.FilteredList).Methods("POST")
	router.HandleFunc("/list", dir.List).Methods("POST")
	router.HandleFunc("/remove", dir.Remove).Methods("POST")
	router.HandleFunc("/removeRecursive", dir.Rmdirs).Methods("POST")
}

func RegisterFileRoutes(router *mux.Router) {
	router.HandleFunc("/compare", file.Compare).Methods("POST")
	router.HandleFunc("/copyUrl", file.CopyURL).Methods("POST")
	router.HandleFunc("/copy", file.Copy).Methods("POST")
	router.HandleFunc("/delete", file.Delete).Methods("POST")
	router.HandleFunc("/moveBackupDir", file.MoveBackupDir).Methods("POST")
	router.HandleFunc("/move", file.Move).Methods("POST")
	router.HandleFunc("/read", file.Read).Methods("POST")
	router.HandleFunc("/write", file.Write).Methods("POST")

	// TODO:
	// router.HandleFunc("/stat", file.Stat).Methods("POST")
	// router.HandleFunc("/rename", file.Rename).Methods("POST")
}

func RegisterSyncRoutes(router *mux.Router) {
	router.HandleFunc("/copy", sync.Copy).Methods("POST")
	router.HandleFunc("/move", sync.Move).Methods("POST")
	router.HandleFunc("/sync", sync.Sync).Methods("POST")
}
