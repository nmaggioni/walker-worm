package patchers

import "fmt"

// Bash is a patcher for shell (bash/sh) files.
func Bash(path string) bool {
	fmt.Printf("%s is a writable Bash file\n", path)
	return true
}
