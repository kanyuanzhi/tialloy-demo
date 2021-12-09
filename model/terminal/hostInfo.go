package terminal

type HostBasicInfo struct {
	Hostname        string `json:"hostname"`
	BootTime        uint64 `json:"bootTime"`
	OS              string `json:"os"`               // ex: freebsd, linux
	Platform        string `json:"platform"`         // ex: ubuntu, linuxmint
	PlatformFamily  string `json:"platform_family"`  // ex: debian, rhel
	PlatformVersion string `json:"platform_version"` // version of the complete OS
	KernelVersion   string `json:"kernel_version"`   // version of the OS kernel (if available)
	KernelArch      string `json:"kernel_arch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
	User            string `json:"user"`
}

func NewHostBasicInfo() *HostBasicInfo {
	return &HostBasicInfo{}
}

type HostRunningInfo struct {
	Uptime uint64 `json:"uptime"`
	Procs  uint64 `json:"procs"` // number of processes
}

func NewHostRunningInfo() *HostRunningInfo {
	return &HostRunningInfo{}
}
