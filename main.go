package main

import (
	"./patchers"
	"./walker"
)

var ignoredPaths = []string{
	"/dev",
	"/proc",
	"/vagrant",
}

// Patchers must return `true` if a file was patched, `false` if not.
var filePatchers = map[string]func(string) bool{
	"py":   patchers.Python,
	"py2":  patchers.Python,
	"py3":  patchers.Python,
	"sh":   patchers.Bash,
	"bash": patchers.Bash,
}

func main() {
	walker.Walk(ignoredPaths, filePatchers)
}
