package main

import (
	"errors"
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

var DefaultContentDir = []string{"html", "htdocs", "content"}

const DefaultHost = "localhost"
const DefaultPort = 22222

func Log(msg string) {
	fmt.Println(msg)
}

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
		Log("Your PC is not supported.")
	}
}

func hasDotPrefix(path string) bool {
	items := strings.Split(path, "/")
	for _, item := range items {
		if strings.HasPrefix(item, ".") {
			return true
		}
	}
	return false
}

func removeCharset(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		t := time.Now()
		Log(t.Format("2006-01-02 15:04:05") + " " + string(r.URL.Path))
		if strings.HasSuffix(r.URL.Path, ".html") {
			header.Del("Content-Type")
			header.Add("Content-Type", "text/html")
		}
		if hasDotPrefix(r.URL.Path) {
			Log(t.Format("2006-01-02 15:04:05") + " Error: not found " + string(r.URL.Path))
			http.NotFound(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}

func findRootDir(path string, dirs []string) (string, error) {
	for _, dir := range dirs {
		rootDir := rootDir(path, dir)
		if _, err := os.Stat(rootDir); err == nil {
			return rootDir, nil
		}
	}
	return "", errors.New("not found")
}

func doMain() int {
	path, err := os.Executable()
	Log("ExecutablePath: "+path)
	if err != nil {
		Log(err.Error())
		return 1
	}
	rootDir, err := findRootDir(path, DefaultContentDir)
	if err != nil {
		Log("コンテンツのディレクトリが見つかりませんでした")
		return 1
	}
	Log("rootDir: "+rootDir)
	http.Handle("/", removeCharset(http.FileServer(http.Dir(rootDir))))

	Log("サーバ起動中...")

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
