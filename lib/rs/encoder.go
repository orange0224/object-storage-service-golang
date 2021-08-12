package rs

import (
	"github.com/klauspost/reedsolomon"
	"io"
)

type encoder struct {
	writers []io.Writer
	encode  reedsolomon.Encoder
	cache   []byte
}

func NewEncoder(writers []io.Writer) *encoder {
	encode, _ := reedsolomon.New(DataShard, ParityShard)
	return &encoder{writers, encode, nil}
}

func (e *encoder) Write(p []byte) (count int, err error) {
	length := len(p)
	current := 0
	for length != 0 {
		next := BlockSie - len(e.cache)
		if next > length {
			next = length
		}
		e.cache = append(e.cache, p[current:current+next]...)
		if len(e.cache) == BlockSie {
			e.Flush()
		}
		current += next
		length -= next
	}
	return len(p), nil
}

func (e *encoder) Flush() {
	if len(e.cache) == 0 {
		return
	}
	shards, _ := e.encode.Split(e.cache)
	e.encode.Encode(shards)
	for i := range shards {
		e.writers[i].Write(shards[i])
	}
	e.cache = []byte{}
}
