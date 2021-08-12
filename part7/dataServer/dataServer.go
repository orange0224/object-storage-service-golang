package main

import (
	"log"
	"net/http"
	"os"
	"storage/part7/dataServer/heartbeat"
	"storage/part7/dataServer/locate"
	"storage/part7/dataServer/objects"
	"storage/part7/dataServer/temp"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
