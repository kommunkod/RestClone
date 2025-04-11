package rclone

import (
	"github.com/rclone/rclone/fs/operations"
)

type RemotePathRequest struct {
	Remote RemoteConfiguration `json:"remote"`
	Path   string              `json:"path"`
}

type ReadFileRequest RemotePathRequest
type DeleteFileRequest RemotePathRequest

type ListFilesRequest struct {
	RemotePathRequest
	Recurse bool                   `json:"recurse"`
	Options operations.ListJSONOpt `json:"options"`
}

type FilterType string

const (
	FilterTypePrefix   FilterType = "prefix"
	FilterTypeSuffix   FilterType = "suffix"
	FilterTypeRegex    FilterType = "regex"
	FilterTypeWildcard FilterType = "wildcard"
)

type FilteredListFilesRequest struct {
	ListFilesRequest
	FilterType FilterType `json:"filterType"`
	Filter     string     `json:"filter"`
}

type WriteFileRequest struct {
	RemotePathRequest
	Overwrite bool   `json:"overwrite"`
	File      []byte `json:"file"`
}

type BulkRenameFilesRequest struct {
	RemotePathRequest
	NameMap map[string]string `json:"nameMap"`
}

type SourceDestinationRequest struct {
	SourceRemote      RemoteConfiguration `json:"sourceRemote"`
	DestinationRemote RemoteConfiguration `json:"destinationRemote"`
	SourcePath        string              `json:"sourcePath"`
	DestinationPath   string              `json:"destinationPath"`
}

type RmdirRequest RemotePathRequest
type RmdirsRequest struct {
	RemotePathRequest
	LeaveRoot bool `json:"leaveRoot"`
}

type CopyFileRequest SourceDestinationRequest
type MoveFileRequest SourceDestinationRequest
type MoveBackupDirRequest SourceDestinationRequest

type CopyURLRequest struct {
	RemotePathRequest
	URL                   string `json:"url"`
	AutoFilename          bool   `json:"autoFilename"`
	DstFilenameFromHeader bool   `json:"dstFilenameFromHeader"`
	NoClobber             bool   `json:"noClobber"`
}

type SyncRequest struct {
	SourceDestinationRequest
	CopyEmptyDirs bool `json:"copyEmptyDirs"`
}

type SyncCopyDirRequest SyncRequest

type SyncMoveDirRequest struct {
	SyncRequest
	DeleteEmptySrcDirs bool `json:"deleteEmptySrcDirs"`
}

type CheckEqualRequest SourceDestinationRequest
