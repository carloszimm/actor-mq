package main

import (
	"crypto/md5"
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/actor-mq/client/actormq"
	"github.com/actor-mq/utils"
)

func getMd5Hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func main() {
	const NumberRequests = 5000

	conn := actormq.Connect(actormq.ConnOptions{"127.0.0.1:8080", "127.0.0.1:8081"})

	conn.CreateChannel("encrypt", utils.PublishSubscribe)
	conn.CreateChannel("encrypted", utils.PublishSubscribe)

	sub := conn.Subscribe("encrypt")
	//defer sub.Unsubscribe()

	for i := 0; i < NumberRequests; i++ {
		msg := <-sub.Ch

		conn.Send("encrypted", []byte(getMd5Hash(string(msg))))
	}

	console.ReadLine()
	sub.Unsubscribe()
}
