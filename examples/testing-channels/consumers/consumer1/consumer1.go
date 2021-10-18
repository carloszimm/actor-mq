package main

import (
	"fmt"

	"github.com/actor-mq/client/actormq"
)

func main() {
	conn := actormq.Connect(actormq.ConnOptions{"127.0.0.1:8080", "127.0.0.1:8081"})

	sub := conn.Subscribe("topicx")

	for {
		msg := <-sub.Ch

		fmt.Printf("onsumer 1, received: %s", string(msg))
	}
}
