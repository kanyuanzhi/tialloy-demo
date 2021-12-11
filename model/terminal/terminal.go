package model

type TerminalBasicPack struct {
	Key  string         `json:"key"` // 本机mac地址
	Data *TerminalBasic `json:"data"`
}

type TerminalBasic struct {
	HostBasic *TerminalHostBasic `json:"host_basic,omitempty" gorm:"embedded;embeddedPrefix:host_"`
	CpuBasic  *TerminalCpuBasic  `json:"cpu_basic,omitempty" gorm:"embedded;embeddedPrefix:cpu_"`
	MemBasic  *TerminalMemBasic  `json:"mem_basic,omitempty" gorm:"embedded;embeddedPrefix:mem_"`
	NetBasic  *TerminalNetBasic  `json:"net_basic,omitempty" gorm:"embedded;embeddedPrefix:net_"`
	DiskBasic *TerminalDiskBasic `json:"disk_basic,omitempty" gorm:"embedded;embeddedPrefix:disk_"`
}

func NewTerminalBasic() *TerminalBasic {
	return &TerminalBasic{
		HostBasic: NewTerminalHostBasic(),
		CpuBasic:  NewTerminalCpuBasic(),
		MemBasic:  NewTerminalMemBasic(),
		NetBasic:  NewTerminalNetBasic(),
		DiskBasic: NewTerminalDiskBasic(),
	}
}

type TerminalRunningPack struct {
	Key  string           `json:"key"`
	Data *TerminalRunning `json:"data"`
}

type TerminalRunning struct {
	HostRunning *TerminalHostRunning `json:"host_running,omitempty"`
	CpuRunning  *TerminalCpuRunning  `json:"cpu_running,omitempty"`
	MemRunning  *TerminalMemRunning  `json:"mem_running,omitempty"`
	NetRunning  *TerminalNetRunning  `json:"net_running,omitempty"`
	DiskRunning *TerminalDiskRunning `json:"disk_running,omitempty"`
}

func NewTerminalRunning() *TerminalRunning {
	return &TerminalRunning{
		HostRunning: NewTerminalHostRunning(),
		CpuRunning:  NewTerminalCpuRunning(),
		MemRunning:  NewTerminalMemRunning(),
		NetRunning:  NewTerminalNetRunning(),
		DiskRunning: NewTerminalDiskRunning(),
	}
}
