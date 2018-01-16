package plugins

import (
    "log"
    "encoding/json"

    "github.com/shirou/gopsutil/disk"
)

type DiskUsage struct {
    Mountpoint string `json:"mountpoint"`
    Fstype string `json:"fstype"`
    Total uint64 `json:"total"`
    Free uint64 `json:"free"`
    Used uint64 `json:"used"`
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
    CheckErr(err)
    for index, disk := range vdisk {
        log.Println("Disk: ", disk.Mountpoint)
        log.Println("INDEX: ", index)
        info := UsageInfo(disk.Mountpoint)
        DiskInformation = append(DiskInformation, info)
    }
    convjson, err := json.Marshal([]DiskUsage(DiskInformation))
    CheckErr(err)

    return convjson
}
