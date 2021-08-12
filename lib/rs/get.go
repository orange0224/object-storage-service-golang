package rs

import (
	"fmt"
	"io"
	"storage/lib/objectStream"
)

type RSGetStream struct {
	*decoder
}

func NewRSGetStream(locateInfo map[int]string, dataServers []string, hash string, size int64) (*RSGetStream, error) {
	if len(locateInfo)+len(dataServers) != AllShard {
		return nil, fmt.Errorf("dataServers number mismatch")
	}
	readers := make([]io.Reader, AllShard)
	for i := 0; i < AllShard; i++ {
		server := locateInfo[i]
		if server == "" {
			locateInfo[i] = dataServers[0]
			dataServers = dataServers[1:]
			continue
		}
		reader, err := objectStream.NewGetStream(server, fmt.Sprintf("%s.%d", hash, i))
		if err == nil {
			readers[i] = reader
		}
	}
	writers := make([]io.Writer, AllShard)
	perShard := (size + DataShard - 1) / DataShard
	var err error
	for i := range readers {
		if readers[i] == nil {
			writers[i], err = objectStream.NewTempPutStream(locateInfo[i], fmt.Sprintf("%s.%d", hash, i), perShard)
			if err != nil {
				return nil, err
			}
		}
	}
	decode := NewDecoder(readers, writers, size)
	return &RSGetStream{decode}, nil
}

func (s *RSGetStream) Close() {
	for i := range s.writers {
		if s.writers[i] != nil {
			s.writers[i].(*objectStream.TempPutStream).Commit(true)
		}
	}
}

func (s *RSGetStream) Seek(offset int64, whence int) (int64, error) {
	if whence != io.SeekCurrent {
		panic("only support SeekCurrent")
	}
	if offset < 0 {
		panic("only support forward seek")
	}
	for offset != 0 {
		length := int64(BlockSie)
		if offset < length {
			length = offset
		}
		buffer := make([]byte, length)
		io.ReadFull(s, buffer)
		offset -= length
	}
	return offset, nil
}
