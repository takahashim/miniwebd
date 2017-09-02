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

func TestHasDotPrefix(t *testing.T)  {
	var pathes = []struct {
		path string
		result bool
	}{
		{"/tmp/foo/bar", false},
		{"/tmp/foo/bar.txt", false},
		{"/tmp/foo/bar.", false},
		{"/.test/foo", true},
		{"/test/.test/foo", true},
		{"/test/...test/foo", true},
		{"/test/foo/.bar", true},
		{"/test/foo/.bar.txt", true},
	}
	for _, tt := range pathes {
		if got, want := hasDotPrefix(tt.path), tt.result; got != want {
			t.Errorf("hasDotPath(): got %v want %v", got, want)
		}
	}
}

