package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/tdarci/prj-999/config"
)

type StaticPathHandler struct {
	*config.Config
	staticPath string
	indexPath  string
}

func NewStaticPathHandler(cfg *config.Config, staticPath string, indexPath string) StaticPathHandler {
	return StaticPathHandler{
		Config:     cfg,
		staticPath: staticPath,
		indexPath:  indexPath,
	}
}

func (h StaticPathHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path, err := filepath.Abs(r.URL.Path)
	h.Logger().Printf("Path is '%s'", path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(h.staticPath, path)

	fm, err := os.Stat(path)
	if os.IsNotExist(err) || fm.IsDir() {
		h.Logger().Printf("serving index.html for '%s'", r.RequestURI)
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Logger().Printf("serving up file: '%s'", path)
	http.ServeFile(w, r, path)
}
