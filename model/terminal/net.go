package model

type TerminalNetBasic struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Mac  string `json:"mac"`
}

func NewTerminalNetBasic() *TerminalNetBasic {
	return &TerminalNetBasic{}
}

type TerminalNetRunning struct {
	BytesSent   uint64 `json:"bytes_sent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytes_recv"`   // number of bytes received
	PacketsSent uint64 `json:"packets_sent"` // number of packets sent
	PacketsRecv uint64 `json:"packets_recv"` // number of packets received
	Errin       uint64 `json:"errin"`        // total number of errors while receiving
	Errout      uint64 `json:"errout"`       // total number of errors while sending
	Dropin      uint64 `json:"dropin"`       // total number of incoming packets which were dropped
	Dropout     uint64 `json:"dropout"`      // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
}

func NewTerminalNetRunning() *TerminalNetRunning {
	return &TerminalNetRunning{}
}
