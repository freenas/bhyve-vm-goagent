package plugins

import (
    "github.com/shirou/gopsutil/cpu"
)

// Returns information about CPU usage.
func CpuInfo() []cpu.TimesStat {
    vcpu, err := cpu.Times(false)
    CheckErr(err)

    return vcpu
}
