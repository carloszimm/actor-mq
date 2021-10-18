package main

import (
	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/actor-mq/broker/invocation"
)

func main() {
	rootContext := actor.EmptyRootContext

	props := actor.PropsFromProducer(func() actor.Actor { return invocation.NewChannelManager() })
	rootContext.SpawnNamed(props, "channelManager")

	console.ReadLine()
}
