package objects

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"storage/lib/utils"
	"storage/part4/dataServer/locate"
)

func getFile(hash string) string {
	filePath := os.Getenv("STORAGE_ROOT") + "/objects/" + hash
	file, _ := os.Open(filePath)
	document := url.PathEscape(utils.CalculateHash(file))
	file.Close()
	fmt.Println("hash:", hash, "-----document: ", document)
	if document != hash {
		log.Println("object hash mismatch,remove,", filePath)
		locate.Del(hash)
		os.Remove(filePath)
		return ""
	}
	return filePath
}
