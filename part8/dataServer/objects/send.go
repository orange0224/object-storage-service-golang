package objects

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func sendFile(w io.Writer, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	gzipStream, err := gzip.NewReader(file)
	if err != nil {
		log.Println(err)
		return
	}
	io.Copy(w, gzipStream)
	gzipStream.Close()
}
