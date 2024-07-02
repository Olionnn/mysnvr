package web

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func StartWeb() {
	fmt.Println("Starting web server...")

	WebRoute()

	srv := &http.Server{
		Handler:      http.DefaultServeMux,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Web server started on http://localhost:8080")
	go openBrowser("http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Web server failed: %v", err)
	}
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}
