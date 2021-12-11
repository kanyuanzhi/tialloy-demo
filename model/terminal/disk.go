package model

type TerminalDiskBasic struct {
	Path   string `json:"path"`
	Fstype string `json:"fstype"`
	Total  uint64 `json:"total"`
}

func NewTerminalDiskBasic() *TerminalDiskBasic {
	return &TerminalDiskBasic{}
}

type TerminalDiskRunning struct {
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func NewTerminalDiskRunning() *TerminalDiskRunning {
	return &TerminalDiskRunning{}
}
