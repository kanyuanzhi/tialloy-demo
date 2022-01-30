package model

type TcpRequest struct {
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}

type TcpCommandRequest struct {
	Command string `json:"command"`
}
