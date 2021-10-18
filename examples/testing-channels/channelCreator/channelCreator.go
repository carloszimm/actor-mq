package main

import (
	console "github.com/AsynkronIT/goconsole"
	"github.com/actor-mq/client/actormq"
	"github.com/actor-mq/utils"
)

func main() {
	conn := actormq.Connect(actormq.ConnOptions{Url: "127.0.0.1:8080"})

	conn.CreateChannel("topicx", utils.PublishSubscribe)
	//conn.CreateChannel("topicx", utils.PointToPoint)

	for text, _ := console.ReadLine(); text != "exit"; text, _ = console.ReadLine() {
		conn.Send("topicx", []byte(text))
	}
}
