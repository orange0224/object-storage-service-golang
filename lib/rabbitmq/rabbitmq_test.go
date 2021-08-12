package rabbitmq

import (
	"encoding/json"
	"testing"
)

const host = "amqp://test:test@192.168.110.134:5672"

func TestPublish(t *testing.T) {
	queue := New(host)
	defer queue.Close()
	queue.Bind("test")

	queue2 := New(host)
	defer queue2.Close()
	queue2.Bind("test")

	queue3 := New(host)
	defer queue.Close()

	expect := "test"
	queue3.Publish("test2", "any")
	queue3.Publish("test", expect)

	consume := queue.Consume()
	message := <-consume
	var actual interface{}
	err := json.Unmarshal(message.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect {
		t.Errorf("expected %s,actual %s", expect, actual)
	}
	if message.ReplyTo != queue3.Name {
		t.Error(message)
	}

	queue2.Send(message.ReplyTo, "test3")
	consume3 := queue3.Consume()
	message = <-consume3
	if string(message.Body) != `"test3"` {
		t.Error(string(message.Body))
	}
}

func TestSend(t *testing.T) {
	queue := New(host)
	defer queue.Close()

	queue2 := New(host)
	defer queue2.Close()

	expect := "test"
	expect2 := "test2"
	queue2.Send(queue.Name, expect)
	queue2.Send(queue2.Name, expect2)

	consume := queue.Consume()
	message := <-consume
	var actual interface{}
	err := json.Unmarshal(message.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect {
		t.Errorf("expected %s,actual %s", expect, actual)
	}
	consume2 := queue2.Consume()
	message = <-consume2
	err = json.Unmarshal(message.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect2 {
		t.Errorf("expected %s,actual %s", expect2, actual)
	}
}
