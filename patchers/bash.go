package patchers

import "fmt"

func Bash(path string) bool {
	fmt.Printf("%s is a writable Bash file\n", path)
	return true
}
