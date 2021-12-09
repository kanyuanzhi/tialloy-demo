package terminal

type MemInfo struct {
	MemBasicInfo   *MemBasicInfo   `json:"mem_basic_info"`
	MemRunningInfo *MemRunningInfo `json:"mem_running_info"`
}

func NewMemInfo() *MemInfo {
	return &MemInfo{
		NewMemBasicInfo(),
		NewMemRunningInfo(),
	}
}

type MemBasicInfo struct {
	Total uint64 `json:"total"`
}

func NewMemBasicInfo() *MemBasicInfo {
	return &MemBasicInfo{}
}

type MemRunningInfo struct {
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func NewMemRunningInfo() *MemRunningInfo {
	return &MemRunningInfo{}
}
