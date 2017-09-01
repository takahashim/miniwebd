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

	removeCharset := func(h http.Handler) http.HandlerFunc {
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

	http.Handle("/", removeCharset(http.FileServer(http.Dir(rootDir))))

	//fmt.Println("RootDir: " + rootDir)
	fmt.Println("aozorahttpd start...")

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

	<-doneCh

	return 0
}

func main() {
	os.Exit(doMain())
}
