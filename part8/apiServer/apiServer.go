package main

import (
	"log"
	"net/http"
	"os"
	"storage/part7/apiServer/heartbeat"
	"storage/part7/apiServer/locate"
	"storage/part7/apiServer/objects"
	"storage/part7/apiServer/temp"
	"storage/part7/apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", version.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
