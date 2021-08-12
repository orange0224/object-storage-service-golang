package locate

import (
	"fmt"
	"os"
	"path/filepath"
	"storage/lib/rabbitmq"
	"strconv"
	"sync"
)

var objects = make(map[string]int)
var mutex sync.Mutex

func Locate(hash string) bool {
	mutex.Lock()
	_, ok := objects[hash]
	mutex.Unlock()
	return ok
}

func Add(hash string) {
	mutex.Lock()
	objects[hash] = 1
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
			fmt.Println(err)
			continue
		}
		exist := Locate(hash)
		if exist {
			queue.Send(message.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}

	}
}

func CollectObject() {
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	for i := range files {
		hash := filepath.Base(files[i])
		objects[hash] = 1
	}
}
