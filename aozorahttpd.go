package main

import (
	"os/exec"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

const AozorabunkoDir = "aozorabunko"

func rootDir(path string) string {
	return filepath.Join(filepath.Dir(path), AozorabunkoDir)
}

func doMain() int {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	rootDir := rootDir(path)
	http.Handle("/", http.FileServer(http.Dir(rootDir)))
	//http.HandleFunc("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(""))))

	fmt.Println("RootDir: " + rootDir)
	fmt.Println("aozorahttdp start")

	doneCh := make(chan error)

	go func() {
		doneCh <- http.ListenAndServe(":22222", nil)
	}()

	url := "http://localhost:22222/index.html"

	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		fmt.Println("Your PC is not supported.")
	}

	<- doneCh

	return 0
}

func main() {
	os.Exit(doMain())
}
