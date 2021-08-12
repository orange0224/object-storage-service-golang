package locate

import (
	"encoding/json"
	"fmt"
	"os"
	"storage/lib/rabbitmq"
	"storage/lib/rs"
	"storage/lib/types"
	"time"
)

func Locate(name string) (locateInfo map[int]string) {
	queue := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	queue.Publish("dataServers", name)
	consume := queue.Consume()
	go func() {
		time.Sleep(time.Second)
		queue.Close()
	}()
	locateInfo = make(map[int]string)
	for i := 0; i < rs.AllShard; i++ {
		message := <-consume
		fmt.Println("receive message:", message.Body)
		if len(message.Body) == 0 {
			return
		}
		var info types.LocateMessage
		json.Unmarshal(message.Body, &info)
		fmt.Println("locate message:", info.Address)
		locateInfo[info.Id] = info.Address
	}
	return
}

func Exist(name string) bool {
	return len(Locate(name)) >= rs.DataShard
}
