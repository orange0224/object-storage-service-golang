package main

import (
	"log"
	"net/http"
	"os"
	"storage/part6/dataServer/heartbeat"
	"storage/part6/dataServer/locate"
	"storage/part6/dataServer/objects"
	"storage/part6/dataServer/temp"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
