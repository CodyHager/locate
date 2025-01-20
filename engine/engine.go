package engine

import (
	"io/fs"
	"strings"

	"locate/util"
)

func SearchForItem(searchingForDir bool, name string, root fs.FS, strict bool, caseSensitive bool, exludes []string) ([]string, error) {
	var paths []string
	util.CreateNewSpinner()
	if err := fs.WalkDir(root, ".", func(path string, d fs.DirEntry, err error) error {
		util.IncrementSpinner(1)
		if (d.IsDir() == searchingForDir) && !shouldExlude(d.Name(), exludes) {
			var filename string
			if caseSensitive {
				filename = d.Name()
			} else {
				name = strings.ToUpper(name)
				filename = strings.ToUpper(d.Name())
			}
			if strict {
				if name == filename {
					paths = append(paths, path)
				}
			} else {
				if strings.Contains(filename, name) {
					paths = append(paths, path)
				}
			}
		}
		return nil
	}); err != nil {
		return paths, err
	} else {
		return paths, nil
	}
}

func shouldExlude(name string, exludes []string) bool {
	for _, e := range exludes {
		if strings.Contains(name, e) {
			return true
		}
	}
	return false
}
