package model

type WebsocketRequestModel struct {
	MsgID uint32   `json:"msg_id,omitempty"`
	Data  []string `json:"data,omitempty"`
}