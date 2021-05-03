package api

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type StaticPathHandler struct {
	StaticPath string
	IndexPath  string
}

func (h StaticPathHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path, err := filepath.Abs(r.URL.Path)
	log.Printf("Path is '%s'", path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(h.StaticPath, path)

	fm, err := os.Stat(path)
	if os.IsNotExist(err) || fm.IsDir() {
		log.Printf("serving index.html for '%s'", r.RequestURI)
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("serving up file: '%s'", path)
	http.ServeFile(w, r, path)
}
