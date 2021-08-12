package main

import (
	"log"
	"net/http"
	"os"
	"storage/part5/apiServer/heartbeat"
	"storage/part5/apiServer/locate"
	"storage/part5/apiServer/objects"
	"storage/part5/apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", version.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))

}
