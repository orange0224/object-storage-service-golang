package main

import (
	"log"
	"net/http"
	"os"
	"storage/part6/apiServer/heartbeat"
	"storage/part6/apiServer/locate"
	"storage/part6/apiServer/objects"
	"storage/part6/apiServer/temp"
	"storage/part6/apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", version.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
