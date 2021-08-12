package rs

import (
	"io"
	"storage/lib/objectStream"
)

type RSResumableGetStream struct {
	*decoder
}

func NewRSResumableGetStream(dataServers, uuids []string, size int64) (*RSResumableGetStream, error) {
	readers := make([]io.Reader, AllShard)
	var err error
	for i := 0; i < AllShard; i++ {
		readers[i], err = objectStream.NewTempGetStream(dataServers[i], uuids[i])
		if err != nil {
			return nil, err
		}
	}
	writers := make([]io.Writer, AllShard)
	decode := NewDecoder(readers, writers, size)
	return &RSResumableGetStream{decode}, nil
}
