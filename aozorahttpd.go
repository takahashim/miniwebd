package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const DefaultContentDir = "content"
const DefaultHost = "localhost"
const DefaultPort = 22222

func rootDir(path, contentDir string) string {
	return filepath.Join(filepath.Dir(path), contentDir)
}

func openBrowser() {
	url := fmt.Sprintf("http://%s:%d/index.html", DefaultHost, DefaultPort)

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
	rootDir := rootDir(path, DefaultContentDir)
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		fmt.Printf("コンテンツのディレクトリが見つかりませんでした\n")
		return 1
	}
	http.Handle("/", removeCharset(http.FileServer(http.Dir(rootDir))))

	//fmt.Println("RootDir: " + rootDir)
	fmt.Println("aozorahttpd start...")

	doneCh := make(chan error)

	go func() {
		doneCh <- http.ListenAndServe(":"+strconv.Itoa(DefaultPort), nil)
	}()

	openBrowser()

	<-doneCh

	return 0
}

func main() {
	os.Exit(doMain())
}
