package model

import (
	"fmt"
	"github.com/Xhofe/alist/conf"
	"golang.org/x/net/webdav"
	"io/fs"
	"time"
)

type File struct {
	Name_     string     `json:"name"`
	Size_     int64      `json:"size"`
	Type      int        `json:"type"`
	Driver    string     `json:"driver"`
	UpdatedAt *time.Time `json:"updated_at"`
	Thumbnail string     `json:"thumbnail"`
	Url       string     `json:"url"`
}

func (f File) Name() string {
	return f.Name_
}

func (f File) Size() int64 {
	return f.Size_
}

func (f File) Mode() fs.FileMode {
	return 0777
}

func (f File) ModTime() time.Time {
	return *f.UpdatedAt
}

func (f File) IsDir() bool {
	return f.Type == conf.FOLDER
}

func (f File) Sys() interface{} {
	return nil
}

var _ fs.FileInfo = (*File)(nil)

func (f File) Close() error {
	return nil
}

func (f File) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (f File) Seek(offset int64, whence int) (int64, error) {
	panic("implement me")
}

func (f File) Readdir(count int) ([]fs.FileInfo, error) {
	panic("implement me")
}

func (f File) Stat() (fs.FileInfo, error) {
	return f,nil
}

func (f File) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("not allow")
}
var _ webdav.File = (*File)(nil)