# RestClone
Stateless Rest API for file operations like read, write, copy, move, sync and more based on rclone. Supports running as a HTTP network service or socket service.

## Supported Backends
All rclone backends are supported, but some are disabled by default to reduce binary size. To enable all backends, uncomment the relevant line in `pkg/rclone/backend.go`.

| Backend | Default Status | Description |
|---------|----------------|-------|
azureblob       | **enabled** | Azure Blob Storage |
azurefiles      | **enabled** | Azure Files |
b2      | **enabled** | Backblaze B2 |
box     | **enabled** | Box |
chunker     | **enabled** | Chunker |
cloudinary      | **enabled** | Cloudinary |
drive       | **enabled** | Google Drive |
dropbox     | **enabled** | Dropbox |
fichier     | disabled | 1Fichier |
filefabric      | **enabled** | FileFabric |
filescom        | disabled | Files.com |
ftp     | **enabled** | FTP |
gofile      | **enabled** | GoFile |
googlecloudstorage      | **enabled** | Google Cloud Storage |
googlephotos        | **enabled** | Google Photos |
hdfs        | **enabled** | HDFS |
hidrive     | **enabled** | HiDrive |
http        | **enabled** | HTTP |
iclouddrive     | **enabled** | iCloud Drive |
imagekit        | **enabled** | ImageKit |
internetarchive     | **enabled** | Internet Archive |
jottacloud      | **enabled** | Jottacloud |
koofr       | **enabled** | Koofr |
linkbox     | **enabled** | Linkbox |
local       | **enabled** | Local |
mailru      | disabled | Mail.ru |
mega        | **enabled** | Mega |
netstorage      | **enabled** | Akamai Netstorage |
onedrive        | **enabled** | OneDrive |
opendrive       | **enabled** | OpenDrive |
oracleobjectstorage     | **enabled** | Oracle Object Storage |
pcloud      | **enabled** | pCloud |
pikpak      | **enabled** | PikPak |
pixeldrain      | **enabled** | PixelDrain |
premiumizeme        | **enabled** | Premiumize.me |
protondrive     | **enabled** | Proton Drive |
putio       | **enabled** | Put.io |
qingstor        | **enabled** | QingStor |
quatrix     | **enabled** | Quatrix |
s3      | **enabled** | S3 |
seafile     | **enabled** | Seafile |
sftp        | **enabled** | SFTP |
sharefile       | **enabled** | ShareFile |
sia     | **enabled** | Sia |
smb     | **enabled** | SMB |
storj       | **enabled** | Storj |
sugarsync       | disabled | SugarSync |
swift       | **enabled** | Swift |
ulozto      | disabled | Ulozto |
uptobox     | disabled | UpToBox |
webdav      | **enabled** | WebDAV |
yandex      | disabled | Yandex |
zoho        | **enabled** | Zoho |
