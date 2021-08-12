package locate

import (
	"fmt"
	"os"
	"storage/lib/rabbitmq"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	queue := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer queue.Close()
	queue.Bind("dataServers")
	consume := queue.Consume()
	for message := range consume {
		object, err := strconv.Unquote(string(message.Body))
		if err != nil {
			panic(err)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			fmt.Println("locate success,prepare send message to ", message.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
			queue.Send(message.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
