package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html")
}

func servePlay(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/play.html")
}

func serveFileManager(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/filemanager.html")
}

func serveStream(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot Load env")
	}
	rtspURI := os.Getenv("rtspURI")

	dir := fmt.Sprintf("live/depan")
	os.MkdirAll(dir, os.ModePerm)

	streamFile := fmt.Sprintf("%s/stream.m3u8", dir)

	cmd := exec.Command("ffmpeg",
		"-rtsp_transport", "tcp",
		"-i", rtspURI,
		"-codec:v", "libx264",
		"-preset", "ultrafast",
		"-g", "50",
		"-f", "hls",
		"-hls_time", "1",
		"-hls_list_size", "3",
		"-hls_flags", "delete_segments",
		streamFile,
	)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("Error creating StderrPipe: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting FFmpeg command: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stderr.Read(buf)
			if err != nil {
				break
			}
			if n > 0 {
				log.Printf("FFmpeg: %s", buf[:n])
			}
		}
	}()

	if err := cmd.Wait(); err != nil {
		log.Printf("FFmpeg command finished with error: %v", err)
	}
}

func serveWebStream(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/stream.html")
}
