package plugins

import (
	"encoding/json"

	"github.com/shirou/gopsutil/host"
)

type HostInfo struct {
	Seconds uint64 `json:"seconds"`
	Days    uint64 `json:"days"`
	Hours   uint64 `json:"hours"`
	Minutes uint64 `json:"minutes"`
}

func Uptime() []byte {
	hostinfo := HostInfo{}
	for i := 0; i < 3; i++ {
		seconds, err := host.Uptime()
		CheckErr(err)
		hostinfo.Seconds = seconds
		hostinfo.Days = hostinfo.Seconds / (60 * 60 * 24)
		hostinfo.Hours = (hostinfo.Seconds - (hostinfo.Days * 60 * 60 * 24)) / (60 * 60)
		hostinfo.Minutes = ((hostinfo.Seconds - (hostinfo.Days * 60 * 60 * 24)) - (hostinfo.Hours * 60 * 60)) / 60
	}
	convjson, err := json.Marshal(hostinfo)
	CheckErr(err)

	return convjson
}
