package types

import (
	"os"
	"time"
)

const (
	FileStatusOK        = "ok"
	FileStatusUploading = "uploading"
	FileStatusDeleted   = "deleted"
)
const (
	FileTypeBin      = "binary"
	FileTypeMarkdown = "markdown"
	FileTypeText     = "text"
)

// Metadata
type FileMetadata struct {
	Name    string            `json:"name"`
	Owner   string            `json:"owner"`
	Project string            `json:"project"`
	Path    string            `json:"path"`
	Size    int64             `json:"size"`
	Hash    string            `json:"hash"`
	Type    string            `json:"type"`
	Mode    os.FileMode       `json:"mode"`
	ModTime time.Time         `json:"modified"`
	CrdTime time.Time         `json:"created"`
	Version string            `json:"version"`
	Status  string            `json:"status"`
	Labels  map[string]string `json:"labels"`
}

type FilePath struct {
	Hash string
	Path string
}
