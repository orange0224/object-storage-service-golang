package main

import (
	"log"
	"net/http"
	"os"
	"storage/part4/dataServer/heartbeat"
	"storage/part4/dataServer/locate"
	"storage/part4/dataServer/objects"
	"storage/part4/dataServer/temp"
)

func main() {
	locate.CollectObject()
	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
