package objects

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"storage/lib/utils"
	"storage/part4/apiServer/locate"
)

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}
	stream, err := putStream(url.PathEscape(hash), size)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	reader := io.TeeReader(r, stream)
	caculate := utils.CalculateHash(reader)
	if caculate != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch,caculated=%s,requestd=%s", caculate, hash)
	}
	stream.Commit(true)
	return http.StatusOK, nil
}
