package main

import (
	"log"
	"net/http"
	"os"
	"storage/part4/apiServer/heartbeat"
	"storage/part4/apiServer/locate"
	"storage/part4/apiServer/objects"
	"storage/part4/apiServer/version"
)

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", version.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))

}
