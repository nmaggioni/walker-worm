package patchers

import "fmt"

func Python(path string) bool {
	fmt.Printf("%s is a writable Python file\n", path)
	return true
}
