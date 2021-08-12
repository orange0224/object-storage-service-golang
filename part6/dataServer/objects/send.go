package objects

import (
	"io"
	"os"
)

func sendFile(w io.Writer, filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	io.Copy(w, file)

}
