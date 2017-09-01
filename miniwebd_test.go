package main

import (
	"testing"
)

func TestRootDir(t *testing.T) {
	var dirs = []struct {
		input, result string
	}{
		{"/tmp/foo/bar", "/tmp/foo/aozorabunko"},
	}
	for _, tt := range dirs {
		if got, want := rootDir(tt.input, "aozorabunko"), tt.result; got != want {
			t.Errorf("rootDir(): got %v want %v", got, want)
		}
	}
}
