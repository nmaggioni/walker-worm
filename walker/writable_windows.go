package walker

import (
	"golang.org/x/sys/unix"
	"os"
)

func writable(path string) bool {
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	defer file.Close()

	return err != nil
}
