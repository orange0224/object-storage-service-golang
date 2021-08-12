package rs

import (
	"fmt"
	"io"
	"storage/lib/objectStream"
)

type RSPutStream struct {
	*encoder
}

func NewRSPutStream(dataServers []string, hash string, size int64) (*RSPutStream, error) {
	if len(dataServers) != AllShard {
		return nil, fmt.Errorf("dataServer number mismatch")
	}
	perShard := (size + DataShard - 1) / DataShard
	writers := make([]io.Writer, AllShard)
	var err error
	for i := range writers {
		writers[i], err = objectStream.NewTempPutStream(dataServers[i], fmt.Sprintf("%s.%d", hash, i), perShard)
		if err != nil {
			return nil, err
		}
	}
	encode := NewEncoder(writers)
	return &RSPutStream{encode}, nil
}

func (s *RSPutStream) Commit(success bool) {
	s.Flush()
	for i := range s.writers {
		s.writers[i].(*objectStream.TempPutStream).Commit(success)
	}
}
