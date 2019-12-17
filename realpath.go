package gokit

import (
	"path/filepath"
	"runtime"
)

// Get file's real path, for example, go test issue
func Realpath(path string) string {
	_, file, _, _ := runtime.Caller(1)
	realpath, _ := filepath.Abs(filepath.Join(filepath.Dir(file), path))
	return realpath
}
