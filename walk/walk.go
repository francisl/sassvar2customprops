package walk

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	cp "cp_migrator/cp_converter"
)


var filesToSearch = regexp.MustCompile("\\w*.s?css")
var excludePaths = [5]*regexp.Regexp{
	regexp.MustCompile("\\*node_modules*"),
	regexp.MustCompile("\\*dist*"),
	regexp.MustCompile("\\*build*"),
	regexp.MustCompile("\\*public*"),
	regexp.MustCompile("\\/\\.\\w*"),
}

func CssConverter(path string) []os.FileInfo {
	excludedPaths := []string{}
	excludedDirs := []string{}
	excludedFiles := []string{}
	convertedFiles := []string{}

	cssFiles := []os.FileInfo{}

	e := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			for i := range excludePaths {
				if excludePaths[i].MatchString(path) {
					excludedPaths = append(excludedPaths, path)
					return nil
				}
			}

			if info.IsDir() {
				excludedDirs = append(excludedDirs, path)
				return nil
			}

			if filesToSearch.MatchString(info.Name()) {
				println("====================================")
				println("Path: ", path)
				println("====================================")
				convertedFiles = append(convertedFiles, path)
				cp.ConvertFile(path)
				cssFiles = append(cssFiles, info)
			} else {
				excludedFiles = append(excludedFiles, path)
			}
		}
		return nil
	})

	if e != nil {
		log.Fatal(e)
	}

	println("Excluded Paths: ", len(excludedPaths))
	println("Excluded Dirs: ", len(excludedDirs))
	println("Excluded Files: ", len(excludedFiles))
	println("Converted Files: ", len(convertedFiles))

	return cssFiles
}
