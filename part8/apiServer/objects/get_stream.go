package objects

import (
	"fmt"
	"storage/lib/rs"
	"storage/part7/apiServer/heartbeat"
	"storage/part7/apiServer/locate"
)

func GetStream(hash string, size int64) (*rs.RSGetStream, error) {
	locateInfo := locate.Locate(hash)
	if len(locateInfo) < rs.DataShard {
		return nil, fmt.Errorf("object %s locate fail,result %v", hash, locateInfo)
	}
	dataServers := make([]string, 0)
	if len(locateInfo) != rs.AllShard {
		dataServers = heartbeat.ChooseRandomDataServers(rs.AllShard-len(locateInfo), locateInfo)
	}
	return rs.NewRSGetStream(locateInfo, dataServers, hash, size)
}
