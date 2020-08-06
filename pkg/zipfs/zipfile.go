package zipfs

import (
	"bytes"
	"os"
)

// ZipFile implements an HTTP File
type ZipFile struct {
	name     string
	fileInfo os.FileInfo
	reader   *bytes.Reader
}

func NewZipFile(n string, f os.FileInfo, r *bytes.Reader) ZipFile {
	return ZipFile {
		name    : n,
		fileInfo: f,
		reader  : r,
	}
}

// Close close the file
func (f ZipFile) Close() error {
	return nil
}

// Read reads the file
func (f ZipFile) Read(p []byte) (int, error) {
	return f.reader.Read(p)
}

// Seek seeks bytes
func (f ZipFile) Seek(offset int64, whence int) (int64, error) {
	return f.reader.Seek(offset, whence)
}

// Readdir func
func (f ZipFile) Readdir(count int) ([]os.FileInfo, error) {
	return []os.FileInfo{ f.fileInfo }, nil
}
// Stat func
func (f ZipFile) Stat() (os.FileInfo, error) {
	return f.fileInfo, nil
}