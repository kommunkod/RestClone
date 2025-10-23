using System.Text.Json.Serialization;

namespace Kommunkod.Restclone;

public abstract class RestcloneRequest<T>
{
    internal abstract string uri { get; set; }
    internal string method { get; set; } = "POST";

    internal Dictionary<string, List<string>> headers = new Dictionary<string, List<string>>
    {
        { "Content-Type", new List<string> { "application/json" } }
    };
    
    internal Interop.RequestFormat GetRequest()
    {
        Console.WriteLine("Method " + this.method);
        Console.WriteLine("URL " + this.uri );
        var reqObj = new Interop.RequestFormat
        {
            Method = this.method,
            Url = this.uri,
            Headers = this.headers
        };
        
        reqObj.SetBody(this);
        return reqObj;
    }

    private Interop.ResponseFormat<T> Invoke()
    {
        var req = this.GetRequest();
        return req.Invoke<T>();
    }

    public T Execute()
    {
        var req = this.Invoke();
        var err = req.IsError();
        if (err != null)
        {
            throw err;
        }

        return req.GetBodyObject();
    }
}

public class RestcloneResponse
{
    
}

public class DirListRequest : RestcloneRequest<DirListResponse>
{
    internal override string uri { get; set; } = "http://restclone/api/v1/dir/list";
    internal string method { get; set; } = "PAST";
    public Remote remote { get; set;  }
    public string path { get; set; }
    public DirListRequestOptions? options { get; set; }
}

public struct DirListRequestOptions
{
    public bool? dirsOnly   { get; set; }
    public bool? filesOnly  { get; set; }
    public List<string>? hashTypes { get; set; }
    public bool? metadata { get; set; }
    public bool? noMimeType { get; set; }
    public bool? noModTime { get; set; }
    public bool? recurse { get; set; }
    public bool? showEncrypted { get; set; }
    public bool? showHash { get; set; }
    public bool? showOrigIDs { get; set; }
}

public class DirListResponse : RestcloneResponse
{
    [JsonPropertyName("files")]
    public FileItem[] Files { get; set; }
}

public struct FileItem
{
    public string Path { get; set; }
    public string Name { get; set; }
    public Int32 Size { get; set; }
    public string MimeType { get; set; }
    public string ModTime { get; set; }
    public bool IsDir { get; set; }
    public bool IsBucket { get; set; }
}

// parent = private path
// child = public fields to be serialized and private path

// public struct ListRequest
// {
//     
// }
//
// public struct ListRequestOptions
// {
//     
// }


public class Operations
{
    public static DirListResponse? DirList(Remote remote, string path, bool recurse = false)
    {
        var request = new DirListRequest()
        {
            remote = remote,
            path = path,
            options = new DirListRequestOptions{
                dirsOnly = false,
                filesOnly = false,
                recurse = recurse,
            }
        };

        return request.Execute();
    }
}



// type RemoteConfiguration struct {
// Name       string                 `json:"name"`
// Type       string                 `json:"type"`
// Parameters map[string]interface{} `json:"parameters"`
// Options    fscfg.UpdateRemoteOpt  `json:"options"`
// }



// type RemotePathRequest struct {
// 	Remote RemoteConfiguration `json:"remote"`
// 	Path   string              `json:"path"`
// }
//
// type ReadFileRequest RemotePathRequest
// type DeleteFileRequest RemotePathRequest
//
// type ListFilesRequest struct {
// 	RemotePathRequest
// 	Recurse bool                   `json:"recurse"`
// 	Options operations.ListJSONOpt `json:"options"`
// }
//
// type FilterType string
//
// const (
// 	FilterTypePrefix   FilterType = "prefix"
// 	FilterTypeSuffix   FilterType = "suffix"
// 	FilterTypeRegex    FilterType = "regex"
// 	FilterTypeWildcard FilterType = "wildcard"
// )
//
// type FilteredListFilesRequest struct {
// 	ListFilesRequest
// 	FilterType FilterType `json:"filterType"`
// 	Filter     string     `json:"filter"`
// }
//
// type WriteFileRequest struct {
// 	RemotePathRequest
// 	Overwrite bool   `json:"overwrite"`
// 	File      []byte `json:"file"`
// }
//
// type BulkRenameFilesRequest struct {
// 	RemotePathRequest
// 	NameMap map[string]string `json:"nameMap"`
// }
//
// type SourceDestinationRequest struct {
// 	SourceRemote      RemoteConfiguration `json:"sourceRemote"`
// 	DestinationRemote RemoteConfiguration `json:"destinationRemote"`
// 	SourcePath        string              `json:"sourcePath"`
// 	DestinationPath   string              `json:"destinationPath"`
// }
//
// type RmdirRequest RemotePathRequest
// type RmdirsRequest struct {
// 	RemotePathRequest
// 	LeaveRoot bool `json:"leaveRoot"`
// }
//
// type CopyFileRequest SourceDestinationRequest
// type MoveFileRequest SourceDestinationRequest
// type MoveBackupDirRequest SourceDestinationRequest
//
// type CopyURLRequest struct {
// 	RemotePathRequest
// 	URL                   string `json:"url"`
// 	AutoFilename          bool   `json:"autoFilename"`
// 	DstFilenameFromHeader bool   `json:"dstFilenameFromHeader"`
// 	NoClobber             bool   `json:"noClobber"`
// }
//
// type SyncRequest struct {
// 	SourceDestinationRequest
// 	CopyEmptyDirs bool `json:"copyEmptyDirs"`
// }
//
// type SyncCopyDirRequest SyncRequest
//
// type SyncMoveDirRequest struct {
// 	SyncRequest
// 	DeleteEmptySrcDirs bool `json:"deleteEmptySrcDirs"`
// }
//
// type CheckEqualRequest SourceDestinationRequest

// func RegisterRoutes(router *mux.Router) {
//     RegisterBulkRoutes(router.PathPrefix("/bulk").Subrouter())
//     RegisterDirectoryRoutes(router.PathPrefix("/dir").Subrouter())
//     RegisterFileRoutes(router.PathPrefix("/file").Subrouter())
//     RegisterSyncRoutes(router.PathPrefix("/sync").Subrouter())
// }
//
// func RegisterBulkRoutes(router *mux.Router) {
//     router.HandleFunc("/rename", bulk.Rename).Methods("POST")
// }
//
// func RegisterDirectoryRoutes(router *mux.Router) {
//     router.HandleFunc("/listFilter", dir.FilteredList).Methods("POST")
//     router.HandleFunc("/list", dir.List).Methods("POST")
//     router.HandleFunc("/remove", dir.Remove).Methods("POST")
//     router.HandleFunc("/removeRecursive", dir.Rmdirs).Methods("POST")
// }
//
// func RegisterFileRoutes(router *mux.Router) {
//     router.HandleFunc("/compare", file.Compare).Methods("POST")
//     router.HandleFunc("/copyUrl", file.CopyURL).Methods("POST")
//     router.HandleFunc("/copy", file.Copy).Methods("POST")
//     router.HandleFunc("/delete", file.Delete).Methods("POST")
//     router.HandleFunc("/moveBackupDir", file.MoveBackupDir).Methods("POST")
//     router.HandleFunc("/move", file.Move).Methods("POST")
//     router.HandleFunc("/read", file.Read).Methods("POST")
//     router.HandleFunc("/write", file.Write).Methods("POST")
//
//     // TODO:
//     // router.HandleFunc("/stat", file.Stat).Methods("POST")
//     // router.HandleFunc("/rename", file.Rename).Methods("POST")
// }
//
// func RegisterSyncRoutes(router *mux.Router) {
//     router.HandleFunc("/copy", sync.Copy).Methods("POST")
//     router.HandleFunc("/move", sync.Move).Methods("POST")
//     router.HandleFunc("/sync", sync.Sync).Methods("POST")
// }


