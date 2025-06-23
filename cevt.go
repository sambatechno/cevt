package cevt

import (
	"github.com/sambatechno/cevt/gen/cevt/msg"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewEmptyUserEvent() *msg.UserEvent {
	return &msg.UserEvent{}
}

// NewUserEvent generates new user event struct with the default fields filled, so only `Body` needs to be filled.
func NewUserEvent() *msg.UserEvent {
	return &msg.UserEvent{
		CreateTimestamp: timestamppb.Now(),
	}
}
