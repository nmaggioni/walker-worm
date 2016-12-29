package walker

import "golang.org/x/sys/unix"

func writable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}
