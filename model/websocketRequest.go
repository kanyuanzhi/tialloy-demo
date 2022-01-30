package model

type WebsocketRequest struct {
	MsgID uint32   `json:"msg_id"`
	Data  []string `json:"data"`
}

type WebsocketCommandRequest struct {
	MsgID   uint32   `json:"msg_id"` // 111
	Data    []string `json:"data"`
	Command string   `json:"command"`
}
