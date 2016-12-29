package walker

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func getFileExtension(filePath string) string {
	extension := path.Ext(filePath)

	if len(extension) > 0 {
		return extension[1:]
	}
	return extension

}

// https://gist.github.com/rayrutjes/db9b9ea8e02255d62ce2
func getMimeType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer[:n])
	if index := strings.Index(contentType, ";"); index != -1 {
		contentType = contentType[:index]
	}

	return contentType, nil
}
