package plugins

import (
	"encoding/json"

	"github.com/shirou/gopsutil/disk"
)

var FSTYPE = map[string]bool{
	"ext3": true,
	"ext4": true,
	"zfs":  true,
	"ufs":  true,
}

type DiskUsage struct {
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Total      uint64 `json:"total"`
	Free       uint64 `json:"free"`
	Used       uint64 `json:"used"`
}

var DiskInformation []DiskUsage

func UsageInfo(path string) DiskUsage {
	usage, err := disk.Usage(path)
	CheckErr(err)

	diskusage := DiskUsage{}
	diskusage.Mountpoint = usage.Path
	diskusage.Fstype = usage.Fstype
	diskusage.Total = usage.Total
	diskusage.Free = usage.Free
	diskusage.Used = usage.Used

	return diskusage
}

// Returns information about all slices.
func DiskInfo() []byte {
	vdisk, err := disk.Partitions(true)
	DiskInformation := []DiskUsage{}
	CheckErr(err)
	for _, disk := range vdisk {
		if FSTYPE[disk.Fstype] {
			info := UsageInfo("/")
			DiskInformation = append(DiskInformation, info)
			break
		}
	}
	convjson, err := json.Marshal([]DiskUsage(DiskInformation))
	CheckErr(err)

	return convjson
}
