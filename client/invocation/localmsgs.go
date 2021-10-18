package invocation

import "github.com/actor-mq/messages"

type SubscribeMsg struct {
	RemoteMsg *messages.SubscribeMsg
	Ch        chan []byte
}
