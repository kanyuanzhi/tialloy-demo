package model

type TerminalCpuBasic struct {
	ModelName     string `json:"model_name"`
	PhysicalCores int    `json:"physical_cores"`
	LogicalCores  int    `json:"logical_cores"`
}

func NewTerminalCpuBasic() *TerminalCpuBasic {
	return &TerminalCpuBasic{}
}

type TerminalCpuRunning struct {
	TotalPercent []float64 `json:"total_percent"`
	PerPercent   []float64 `json:"per_percent"`
}

func NewTerminalCpuRunning() *TerminalCpuRunning {
	return &TerminalCpuRunning{}
}
