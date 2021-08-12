package main

import (
	"log"
	"net/http"
	"os"
	"storage/part5/dataServer/heartbeat"
	"storage/part5/dataServer/locate"
	"storage/part5/dataServer/objects"
	"storage/part5/dataServer/temp"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
