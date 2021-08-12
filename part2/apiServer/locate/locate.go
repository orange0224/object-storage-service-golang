package locate

import (
	"os"
	"storage/lib/rabbitmq"
	"strconv"
	"time"
)

func Locate(name string) string {
	queue := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	queue.Publish("dataServers", name)
	consume := queue.Consume()
	go func() {
		time.Sleep(time.Second)
		queue.Close()
	}()
	message := <-consume
	str, _ := strconv.Unquote(string(message.Body))
	return str
}

func Exist(name string) bool {
	return Locate(name) != ""
}
