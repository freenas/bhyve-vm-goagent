package plugins

import (
    "github.com/shirou/gopsutil/net"
)

// Returns information about IFACE settings.
func NetInfo() []net.InterfaceStat {
    vnet, err := net.Interfaces()
    CheckErr(err)

    return vnet
}
