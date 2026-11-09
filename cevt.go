package cevt

import (
	"github.com/sambatechno/cevt/gen/cevt/msg"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Verification methods for the AccountVerified user event (INT-492). The proto
// field AccountVerified.method is a free string, so producers (user-service) and
// consumers (loyalty-service, integration-service) must share these exact values
// — reference these constants rather than typing the literals, so a typo can
// never silently fail a match.
const (
	VerificationMethodEmailOTP = "EMAIL_OTP"
	VerificationMethodPhoneOTP = "PHONE_OTP"
	VerificationMethodSocial   = "SOCIAL"
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
