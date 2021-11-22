package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
	"net/http"
	"os"
	"path/filepath"
)

func WebDav(r *gin.Engine) {
	dav := r.Group("/dav")
	dav.Any("", Dav)
	dav.Any("/*path", Dav)
}

var davFS = webdav.Handler{
	Prefix:     "/dav",
	FileSystem: AListFileSystem{},
	LockSystem: nil,
}

func Dav(c *gin.Context) {
	r := c.Request
	w := c.Writer
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// validate username and password
	if username != "user" || password != "123456" {
		http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
		http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
		return
	}
	davFS.ServeHTTP(w, r)
}

type AListFileSystem struct {
}

var ReadOnlyErr = fmt.Errorf("WebDAV: Read Only")

func (A AListFileSystem) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	return ReadOnlyErr
}

func (A AListFileSystem) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	return nil, nil
}

func (A AListFileSystem) RemoveAll(ctx context.Context, name string) error {
	return ReadOnlyErr
}

func (A AListFileSystem) Rename(ctx context.Context, oldName, newName string) error {
	return ReadOnlyErr
}

func (A AListFileSystem) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	dir := filepath.Dir(name)
	name = filepath.Base(name)
	account, path, driver, err := ParsePath(dir)
	if err != nil {
		return nil, err
	}
	_, files, err := driver.Path(path, account)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.Name() == name {
			return file, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

var _ webdav.FileSystem = (*AListFileSystem)(nil)
