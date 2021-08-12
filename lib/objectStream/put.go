package objectStream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	channel := make(chan error)
	go func() {
		request, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		client := http.Client{}
		result, err := client.Do(request)
		if err == nil && result.StatusCode != http.StatusOK {
			err = fmt.Errorf("dataServer return http code %d", result.StatusCode)
		}
		channel <- err
	}()
	return &PutStream{writer: writer, c: channel}
}

func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c
}
