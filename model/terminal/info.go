package terminal

type Info struct {
	Key      string    `json:"key,omitempty"`
	InfoData *InfoData `json:"data,omitempty"`
}

type InfoData struct {
	BasicInfo   *BasicInfo   `json:"basic_info,omitempty"`
	RunningInfo *RunningInfo `json:"running_info,omitempty"`
}

type BasicInfo struct {
	Key           string         `json:"key,omitempty"`
	basicInfoData *BasicInfoData `json:"data,omitempty"`
}

type BasicInfoData struct {
	HostBasicInfo *HostBasicInfo `json:"host_basic_info"`
	CpuBasicInfo  *CpuBasicInfo  `json:"cpu_basic_info"`
	MemBasicInfo  *MemBasicInfo  `json:"mem_basic_info"`
	NetBasicInfo  *NetBasicInfo  `json:"net_basic_info"`
	DiskBasicInfo *DiskBasicInfo `json:"disk_basic_info"`
}

type RunningInfo struct {
	Key             string           `json:"key,omitempty"`
	runningInfoData *RunningInfoData `json:"data,omitempty"`
}

type RunningInfoData struct {
	HostRunningInfo *HostRunningInfo `json:"host_running_info"`
	CpuRunningInfo  *CpuRunningInfo  `json:"cpu_running_info"`
	MemRunningInfo  *MemRunningInfo  `json:"mem_running_info"`
	NetRunningInfo  *NetRunningInfo  `json:"net_running_info"`
	DiskRunningInfo *DiskRunningInfo `json:"disk_running_info"`
}
