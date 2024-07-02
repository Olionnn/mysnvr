package web

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func apiFileList(w http.ResponseWriter, r *http.Request) {
	recordingsDir := "recordings"
	var files []string

	err := filepath.Walk(recordingsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp4") {
			relativePath, err := filepath.Rel(recordingsDir, path)
			if err != nil {
				return err
			}
			files = append(files, relativePath)
		}
		return nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
