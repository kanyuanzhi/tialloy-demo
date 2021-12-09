package model

type WebsocketRequest struct {
	MsgID uint32   `json:"msg_id"`
	Data  []string `json:"data"`
}
