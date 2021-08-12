package temp

import (
	"compress/gzip"
	"io"
	"net/url"
	"os"
	"storage/lib/utils"
	"storage/part8/dataServer/locate"
	"strconv"
	"strings"
)

func (t *tempInfo) hash() string {
	str := strings.Split(t.Name, ".")
	return str[0]
}

func (t *tempInfo) id() int {
	str := strings.Split(t.Name, ".")
	id, _ := strconv.Atoi(str[1])
	return id
}

func commitTempObject(dataFile string, info *tempInfo) {
	file, _ := os.Open(dataFile)
	defer file.Close()
	calculated := url.PathEscape(utils.CalculateHash(file))
	file.Seek(0, io.SeekStart)
	path, _ := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + info.Name + "." + calculated)
	writer := gzip.NewWriter(path)
	io.Copy(writer, file)
	writer.Close()
	os.Remove(dataFile)
	locate.Add(info.hash(), info.id())

}
