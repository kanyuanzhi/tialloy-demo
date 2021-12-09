package terminal

type DiskInfo struct {
	DiskBasicInfo   *DiskBasicInfo   `json:"disk_basic_info"`
	DiskRunningInfo *DiskRunningInfo `json:"disk_running_info"`
}

func NewDiskInfo() *DiskInfo {
	return &DiskInfo{
		DiskBasicInfo:   NewDiskBasicInfo(),
		DiskRunningInfo: NewDiskRunningInfo(),
	}
}

type DiskBasicInfo struct {
	Path   string `json:"path"`
	Fstype string `json:"fstype"`
	Total  uint64 `json:"total"`
}

func NewDiskBasicInfo() *DiskBasicInfo {
	return &DiskBasicInfo{}
}

type DiskRunningInfo struct {
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func NewDiskRunningInfo() *DiskRunningInfo {
	return &DiskRunningInfo{}
}
