package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const AozorabunkoDir = "aozorabunko"
const OpenUrl = "http://localhost:22222/index.html"

func rootDir(path string) string {
	return filepath.Join(filepath.Dir(path), AozorabunkoDir)
}

func openBrowser() {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", OpenUrl).Start()
	case "windows":
		exec.Command("rundll32", "OpenUrl.dll,FileProtocolHandler", OpenUrl).Start()
	case "darwin":
		exec.Command("open", OpenUrl).Start()
	default:
		fmt.Println("Your PC is not supported.")
	}
}

func removeCharset(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		fmt.Println(string(r.URL.Path))
		if strings.HasSuffix(r.URL.Path, ".html") {
			header.Del("Content-Type")
			header.Add("Content-Type", "text/html")
		}
		h.ServeHTTP(w, r)
	}
}

func doMain() int {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	rootDir := rootDir(path)

	http.Handle("/", removeCharset(http.FileServer(http.Dir(rootDir))))

	//fmt.Println("RootDir: " + rootDir)
	fmt.Println("aozorahttpd start...")

	doneCh := make(chan error)

	go func() {
		doneCh <- http.ListenAndServe(":22222", nil)
	}()

	openBrowser()

	<-doneCh

	return 0
}

func main() {
	os.Exit(doMain())
}
