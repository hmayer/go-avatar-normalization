package main

import (
	"fmt"
	"go-avatar-normalization/handlers"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func init() {
	clearCachedFiles()
}

func clearCachedFiles() {
	retention := time.Now().Add(-1 * time.Hour)
	files, _ := filepath.Glob(filepath.Join("./resources/images", "*"))
	for _, f := range files {
		info, _ := os.Stat(f)
		if info.ModTime().Before(retention) {
			fmt.Println("Removing old cached file", info.Name())
			os.Remove(f)
		}
	}
}

func clearLoop() {
	for {
		time.Sleep(10 * time.Minute)
		clearCachedFiles()
	}
}

func main() {
	http.HandleFunc("/avatar", handlers.AvatarUploadHandler)
	fmt.Println("Routes defined")
	go clearLoop()
	fmt.Println(http.ListenAndServe(":8000", nil))
}
