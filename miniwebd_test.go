package main

import (
	"testing"
)

func TestRootDir(t *testing.T) {
	var dirs = []struct {
		input, dirname, result string
	}{
		{"/tmp/foo/bar", "content", "/tmp/foo/content"},
		{"/", "content", "/content"},
		{"/tmp/", "content", "/tmp/content"},
	}
	for _, tt := range dirs {
		if got, want := rootDir(tt.input, tt.dirname), tt.result; got != want {
			t.Errorf("rootDir(): got %v want %v", got, want)
		}
	}
}
