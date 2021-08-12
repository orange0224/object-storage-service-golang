package objects

import (
	"fmt"
	"storage/lib/objectStream"
	"storage/part4/apiServer/heartbeat"
)

func putStream(hash string, size int64) (*objectStream.TempPutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectStream.NewTempPutStream(server, hash, size)
}
