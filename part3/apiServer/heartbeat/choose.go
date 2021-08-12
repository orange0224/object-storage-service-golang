package heartbeat

import (
	"fmt"
	"math/rand"
)

func ChooseRandomDataServer() string {
	dataServers := GetDataServers()
	count := len(dataServers)
	fmt.Println("dataServer.size=", len(dataServers))
	if count == 0 {
		return ""
	}
	return dataServers[rand.Intn(count)]
}
