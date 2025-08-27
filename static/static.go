package static

import (
	"os"
	"path/filepath"
	"runtime"
)

func ProjectRoot() string {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	for dir != "/" {
		if _, err := filepath.Abs(filepath.Join(dir, "go.mod")); err == nil {
			if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
				return dir
			}
		}
		dir = filepath.Dir(dir)
	}
	return filepath.Dir(filepath.Dir(filename))
}

func File(subpath string) string {
	return filepath.Join(ProjectRoot(), "static", subpath)
}