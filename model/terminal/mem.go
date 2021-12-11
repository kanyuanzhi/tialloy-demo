package model

type TerminalMemBasic struct {
	Total uint64 `json:"total"`
}

func NewTerminalMemBasic() *TerminalMemBasic {
	return &TerminalMemBasic{}
}

type TerminalMemRunning struct {
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func NewTerminalMemRunning() *TerminalMemRunning {
	return &TerminalMemRunning{}
}
