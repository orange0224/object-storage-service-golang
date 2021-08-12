package locate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"storage/lib/rabbitmq"
	"storage/lib/types"
	"strconv"
	"strings"
	"sync"
)

var objects = make(map[string]int)
var mutex sync.Mutex

func Locate(hash string) int {
	mutex.Lock()
	id, ok := objects[hash]
	mutex.Unlock()
	if !ok {
		return -1
	}
	return id
}

func Add(hash string, id int) {
	mutex.Lock()
	objects[hash] = id
	mutex.Unlock()
}

func Del(hash string) {
	mutex.Lock()
	delete(objects, hash)
	mutex.Unlock()
}

func StartLocate() {
	queue := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer queue.Close()
	queue.Bind("dataServers")
	consume := queue.Consume()
	for message := range consume {
		hash, err := strconv.Unquote(string(message.Body))
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("prepare locate:", hash)
		id := Locate(hash)
		if id != -1 {
			queue.Send(message.ReplyTo, types.LocateMessage{Address: os.Getenv("LISTEN_ADDRESS"), Id: id})
		}
	}
}

func CollectObjects() {
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	for i := range files {
		file := strings.Split(filepath.Base(files[i]), ".")
		if len(file) != 3 {
			panic(files[i])
		}
		hash := file[0]
		id, err := strconv.Atoi(file[1])
		if err != nil {
			panic(err)
		}
		objects[hash] = id
	}
}
