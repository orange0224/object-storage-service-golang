package temp

import (
	"os"
	"storage/part4/dataServer/locate"
)

func commitTempObject(dataFile string, tempInfo *tempInfo) {
	os.Rename(dataFile, os.Getenv("STORAGE_ROOT")+"/objects/"+tempInfo.Name)
	locate.Add(tempInfo.Name)
}
