package objects

import (
	"fmt"
	"storage/lib/objectStream"
	"storage/part2/apiServer/heartbeat"
)

func putStream(object string) (*objectStream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectStream.NewPutStream(server, object), nil
}
