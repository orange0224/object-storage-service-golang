package temp

import (
	"net/url"
	"os"
	"storage/lib/utils"
	"storage/part5/dataServer/locate"
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
	calculated := url.PathEscape(utils.CalculateHash(file))
	file.Close()
	os.Rename(dataFile, os.Getenv("STORAGE_ROOT")+"/objects/"+info.Name+"."+calculated)
	locate.Add(info.hash(), info.id())

}
