package terminal

type CpuBasicInfo struct {
	ModelName     string `json:"model_name"`
	PhysicalCores int    `json:"physical_cores"`
	LogicalCores  int    `json:"logical_cores"`
}

func NewCpuBasicInfo() *CpuBasicInfo {
	return &CpuBasicInfo{}
}

type CpuRunningInfo struct {
	TotalPercent []float64 `json:"total_percent"`
	PerPercent   []float64 `json:"per_percent"`
}

func NewCpuRunningInfo() *CpuRunningInfo {
	return &CpuRunningInfo{}
}
