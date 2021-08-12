package main

import (
	"log"
	"storage/lib/es/es7"
)

const MinVersionCount = 5

func main() {
	buckets, err := es7.SearchVersionStatus(MinVersionCount + 1)
	if err != nil {
		log.Println(err)
		return
	}
	for i := range buckets {
		bucket := buckets[i]
		for v := 0; v < bucket.DocCount-MinVersionCount; v++ {
			es7.DelMetadata(bucket.Key, v+int(bucket.MinVersion.Value))
		}
	}
}
