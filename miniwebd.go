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
		fmt.Printf("%s %s\n", t.Format("2006-01-02 15:04:05"), string(r.URL.Path))
		if strings.HasSuffix(r.URL.Path, ".html") {
			header.Del("Content-Type")
			header.Add("Content-Type", "text/html")
		}
		if hasDotPrefix(r.URL.Path) {
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
			return dir, nil
		}
	}
	return "", errors.New("not found")
}

func doMain() int {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	rootDir, err := findRootDir(path, DefaultContentDir)
	if err != nil {
		fmt.Printf("コンテンツのディレクトリが見つかりませんでした\n")
		return 1
	}
	http.Handle("/", removeCharset(http.FileServer(http.Dir(rootDir))))

	fmt.Println("サーバ起動中...")

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
