package web

import "net/http"

func WebRoute() {

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/play", servePlay)
	http.HandleFunc("/filemanager", serveFileManager)
	http.HandleFunc("/stream", serveWebStream)

	http.HandleFunc("/api/filelist", apiFileList)

	http.Handle("/recordings/", http.StripPrefix("/recordings/", http.FileServer(http.Dir("recordings"))))
	http.Handle("/live/", http.StripPrefix("/live/", http.FileServer(http.Dir("live"))))

	go func() {
		for {
			serveStream(nil, nil)
		}
	}()
}
