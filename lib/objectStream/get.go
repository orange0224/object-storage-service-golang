package objectStream

import (
	"fmt"
	"io"
	"net/http"
)

type GetStream struct {
	reader io.Reader
}

func newGetStream(url string) (*GetStream, error) {
	fmt.Println("rpc:get stream from:", url)
	result, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if result.StatusCode != http.StatusOK {
		fmt.Println("URL:", url)
		return nil, fmt.Errorf("dataServer return http code %d", result.StatusCode)
	}
	return &GetStream{result.Body}, nil
}

func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, object)
	}
	return newGetStream("http://" + server + "/objects/" + object)
}

func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
