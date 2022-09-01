package server

import (
	"path/filepath"

	. "github.com/lyft/goruntime/loader"
)

type CustomDirectoryRefresher struct {
	currDir  string
	watchOps map[FileSystemOp]struct{}
}

var defaultFileSystemOps = map[FileSystemOp]struct{}{
	Write:  {},
	Create: {},
	Chmod:  {},
}

func (d *CustomDirectoryRefresher) WatchDirectory(runtimePath string, appDirPath string) string {
	d.currDir = filepath.Join(runtimePath, appDirPath)
	return d.currDir
}

func (d *CustomDirectoryRefresher) ShouldRefresh(path string, op FileSystemOp) bool {
	var watchOps *map[FileSystemOp]struct{}

	if d.watchOps == nil {
		watchOps = &defaultFileSystemOps
	} else {
		watchOps = &d.watchOps
	}

	if _, opMatches := (*watchOps)[op]; opMatches && filepath.Dir(path) == d.currDir {
		return true
	}
	return false
}
