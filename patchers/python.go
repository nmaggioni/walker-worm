package patchers

import "fmt"

// Python is a patcher for python (py/py2/py3) files.
func Python(path string) bool {
	fmt.Printf("%s is a writable Python file\n", path)
	return true
}
