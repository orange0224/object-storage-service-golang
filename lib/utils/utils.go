package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetOffsetFromHeader(header http.Header) int64 {
	byteRange := header.Get("range")
	if len(byteRange) < 7 {

		return 0
	}
	if byteRange[:6] != "bytes=" {
		return 0
	}
	bytePos := strings.Split(byteRange[6:], "-")
	offset, _ := strconv.ParseInt(bytePos[0], 0, 64)
	return offset
}

func GetHashFromHeader(header http.Header) string {
	digest := header.Get("digest")
	if len(digest) < 9 {
		return ""
	}
	if digest[:8] != "SHA-256=" {
		return ""
	}
	return digest[8:]
}

func GetSizeFromHeader(header http.Header) int64 {
	size, _ := strconv.ParseInt(header.Get("content-length"), 0, 64)
	return size
}

func CalculateHash(reader io.Reader) string {
	hash := sha256.New()
	io.Copy(hash, reader)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
