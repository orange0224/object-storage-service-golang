package main

import (
	"log"
	"net/http"
	"os"
	"storage/part3/dataServer/heartbeat"
	"storage/part3/dataServer/locate"
	"storage/part3/dataServer/objects"
)

func main() {
	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
