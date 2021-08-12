package objects

import (
	"fmt"
	"io"
	"storage/lib/objectStream"
	"storage/part4/apiServer/locate"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectStream.NewGetStream(server, object)
}
