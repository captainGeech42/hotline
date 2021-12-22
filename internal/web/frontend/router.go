package frontend

import (
	"embed"
	"io/fs"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

// embed FS for the react SPA code
//go:embed spa/build/*
var content embed.FS

// implements fs.FS to transparently strip the on-disk file path
// exposed by the embed.FS implementation
// e.g., when a request comes in for /index.html, it needs to be
// convereted to /spa/build/index.html before being read from
// the embed fs
type spaFS struct {
	InnerFS *embed.FS
}

// fix the path when referencing the embedfs
// see frontend.spaFS
func (fs *spaFS) Open(name string) (fs.File, error) {
	fixedPath := path.Join("spa", "build", name)

	return fs.InnerFS.Open(fixedPath)
}

func ConfigureRouter(router *mux.Router) {
	myFS := &spaFS{InnerFS: &content}
	httpFS := http.FileServer(http.FS(myFS))

	router.PathPrefix("/").Handler(httpFS)
}
