package main

import (
	"fmt"
	"log"
	"os"

	"github.com/araujobsd/bhyve-vm-goagent/plugins"
	"github.com/araujobsd/bhyve-vm-goagent/termios"
)

const DEBUG int = 1

var (
	guestInfo    []byte
	vconsolePath = [...]string{
		"/dev/virtio-ports/org.freenas.bhyve-agent",
		"/dev/vtcon/org.freenas.bhyve-agent",
		"/dev/ttyV0.0",
	}
)

func checkConsole() string {
	for _, device := range vconsolePath {
		_, err := os.Stat(device)
		if os.IsNotExist(err) {
			continue
		} else {
			if DEBUG == 1 {
				log.Println("==> Device: ", device)
			}
			return device
		}
	}
	return ""
}

func Run() {
	var dump []byte
	vconsole := checkConsole()
	callback := termios.Read(vconsole)
	if len(callback) > 0 {
		switch opt := string(callback); opt {
		case "mem":
			memoryinfo := plugins.Memory()
			dump = append([]byte(fmt.Sprintf("%v", memoryinfo)))
		case "cpu":
			cpuinfo := plugins.CpuInfo()
			dump = append([]byte(fmt.Sprintf("%v", cpuinfo)))
		case "iface":
			ifaceinfo := plugins.NetInfo()
			dump = append([]byte(fmt.Sprintf("%v", ifaceinfo)))
		case "disk":
			diskinfo := plugins.DiskInfo()
			dump = append(diskinfo)
		case "uptime":
			uptimeinfo := plugins.Uptime()
			dump = append(uptimeinfo)
		default:
			dump = append([]byte("pong"))
		}
		termios.Write(vconsole, dump)
	}
}

func main() {
	for {
		Run()
	}
}
