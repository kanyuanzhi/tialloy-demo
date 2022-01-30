package model

type WebsocketResponse struct {
	MsgID uint32      `json:"msg_id"`
	Data  interface{} `json:"data"`
}
