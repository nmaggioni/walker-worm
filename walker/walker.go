package walker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var _ignoredPaths []string
var _filePatchers map[string]func(string) bool
var wg sync.WaitGroup
var ch chan bool
var walkedPaths uint64
var patchedFiles int

func walkPath(path string, f os.FileInfo, err error) error {
	{
		for _, ignore := range _ignoredPaths {
			if strings.HasPrefix(path, ignore) {
				return nil
			}
		}

		if err != nil {
			if f.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
	}

	if !f.IsDir() && writable(path) {
		if patcher, isPresent := _filePatchers[getFileExtension(path)]; isPresent {
			wg.Add(1)
			go func(path string, ch chan bool) {
				defer wg.Done()
				ch <- patcher(path)
			}(path, ch)
		}
	}

	walkedPaths++
	fmt.Printf("=== Walked %d paths\r", walkedPaths)
	return nil
}

// Walk is the main body of the kit: it traverses the filesystem invoking `walkPath` on each path that is found.
func Walk(ignoredPaths []string, filePatchers map[string]func(string) bool) {
	_ignoredPaths = ignoredPaths
	_filePatchers = filePatchers
	ch = make(chan bool, 100)
	go func(ch chan bool) {
		for {
			didPatch, isOpen := <-ch
			if !isOpen {
				break
			}
			if didPatch {
				patchedFiles++
			}
		}
	}(ch)

	fmt.Println("=== Walker is walking")
	err := filepath.Walk(filesystemRoot, walkPath)

	var retVal string
	if err != nil {
		retVal = err.Error()
	} else {
		retVal = "no errors"
	}

	wg.Wait()
	close(ch)

	fmt.Printf("=== Walker stopped walking: %s\n", retVal)
	fmt.Printf("=== Patched files: %d\n", patchedFiles)
}
