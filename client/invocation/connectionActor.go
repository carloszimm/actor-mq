package invocation

import (
	"log"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/actor-mq/messages"
)

type ConnectionActor struct {
	RemotePid     *actor.PID
	Subscriptions map[string]chan []byte
}

func (state *ConnectionActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.CreateChannelMsg:
		log.Println("[INFO] Received Create Channel Message")

		context.Send(state.RemotePid, msg)

	case *messages.PublishMsg:
		log.Println("[INFO] Received Publish Msg")

		context.Send(state.RemotePid, msg)
	case *messages.NotifyMsg:
		log.Println("[INFO] Received Notify Msg")

		if ch, ok := state.Subscriptions[msg.ChannelName]; ok {
			ch <- msg.Content
		}

	case *SubscribeMsg:
		log.Println("[INFO] Received Subscribe Msg")

		state.Subscriptions[msg.RemoteMsg.ChannelName] = msg.Ch
		context.Send(state.RemotePid, msg.RemoteMsg)

	case *messages.UnsubscribeMsg:
		log.Println("[INFO] Received Unsubscribe Msg")

		delete(state.Subscriptions, msg.ChannelName)
		context.Send(state.RemotePid, msg)
	}
}
