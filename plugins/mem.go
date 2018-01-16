package plugins

import (
    "github.com/shirou/gopsutil/mem"
)

// Returns information about memory usage.
func Memory() *mem.VirtualMemoryStat {
    vm, err := mem.VirtualMemory()
    CheckErr(err)

    return vm
}
