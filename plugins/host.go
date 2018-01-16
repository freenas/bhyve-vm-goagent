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
	seconds, err := host.Uptime()
	CheckErr(err)

	hostinfo.Seconds = seconds
	hostinfo.Days = (seconds / (60 * 60 * 24))
	hostinfo.Hours = (seconds - (hostinfo.Days * 60 * 60 * 24)) / (60 * 60)
	hostinfo.Minutes = ((seconds - (hostinfo.Days * 60 * 60 * 24)) - (hostinfo.Hours * 60 * 60)) / 60
	convjson, err := json.Marshal(hostinfo)

	return convjson
}
