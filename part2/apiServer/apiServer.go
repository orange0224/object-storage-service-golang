package main

import (
	"log"
	"net/http"
	"os"
	"storage/part2/apiServer/heartbeat"
	"storage/part2/apiServer/locate"
	"storage/part2/apiServer/objects"
)

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
