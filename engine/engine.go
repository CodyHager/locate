package engine

import (
	"io/fs"
)

func SearchForFile(name string, root fs.FS) ([]string, error) {
	var paths []string
	if err := fs.WalkDir(root, ".", func(path string, d fs.DirEntry, err error) error {
		// process each item in the root here
		if name == d.Name() {
			paths = append(paths, path)
		}
		return nil
	}); err != nil {
		return paths, err
	} else {
		return paths, nil
	}
}
