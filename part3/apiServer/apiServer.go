package main

import (
	"log"
	"net/http"
	"os"
	"storage/part3/apiServer/heartbeat"
	"storage/part3/apiServer/locate"
	"storage/part3/apiServer/objects"
	"storage/part3/apiServer/version"
)

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", version.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))

}
