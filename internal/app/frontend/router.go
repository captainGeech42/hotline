package frontend

import (
	"embed"
	"io/fs"
	"log"
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

func (fs *spaFS) Open(name string) (fs.File, error) {
	log.Println("called")

	fixedPath := path.Join("spa", "build", name)
	println(fixedPath)

	log.Println("opening", fixedPath)

	return fs.InnerFS.Open(fixedPath)
}

func ConfigureRouter(router *mux.Router) {
	_, err := content.Open("/spa/build/index.html")
	log.Println(err)

	myFS := &spaFS{InnerFS: &content}

	router.Handle("/", http.FileServer(http.FS(myFS)))
}
