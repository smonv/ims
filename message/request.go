package message

import "encoding/json"

// ActionType
const (
	_ = iota
	ActionTypeCreateImage
	ActionTypeCreateTag
)

// Request message
type Request struct {
	ID         string           `json:"id"`
	ActionType int8             `json:"action_type"`
	ActionData *json.RawMessage `json:"action_data"`
}
