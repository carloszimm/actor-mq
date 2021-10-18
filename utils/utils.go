package utils

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/actor-mq/messages"
)

type ChannelType int32

const (
	PublishSubscribe ChannelType = iota
	PointToPoint
)

func NewCreateChannelMsg(channelName string, channelType ChannelType) *messages.CreateChannelMsg {
	return &messages.CreateChannelMsg{channelName, int32(channelType)}
}

func NewPublishMsg(channelName string, content []byte) *messages.PublishMsg {
	return &messages.PublishMsg{channelName, content}
}

func NewNotifyMsg(channelName string, content []byte) *messages.NotifyMsg {
	return &messages.NotifyMsg{channelName, content}
}

func NewSubscribeMsg(channelName string, subscriber *actor.PID) *messages.SubscribeMsg {
	return &messages.SubscribeMsg{channelName, subscriber}
}

func NewUnsubscribeMsg(channelName string, subscriber *actor.PID) *messages.UnsubscribeMsg {
	return &messages.UnsubscribeMsg{channelName, subscriber}
}
