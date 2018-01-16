package main

import (
    "os"
    "fmt"
    "log"

    "github.com/pkg/term"
    "github.com/araujobsd/bhyve-vm-goagent/plugins"
)

const DEBUG int = 0

var (
    guestInfo []byte
    vconsolePath = [...]string {
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

func writeConsole(opt string) {
    vconsole := checkConsole()

    virtiorw, err := term.Open(vconsole, term.Speed(9600), term.RawMode)
    plugins.CheckErr(err)
    virtiorw.Flush()

    if DEBUG == 1 {
        log.Println("==> Opt: ", opt)
    }

    switch opt {
    case "mem":
        meminfo := plugins.Memory()
        guestInfo = append([]byte(fmt.Sprintf("%v", meminfo)))
    case "cpu":
        cpuinfo := plugins.CpuInfo()
        guestInfo = append([]byte(fmt.Sprintf("%v", cpuinfo)))
    case "iface":
        ifaceinfo := plugins.NetInfo()
        guestInfo = append([]byte(fmt.Sprintf("%v", ifaceinfo)))
    default:
        guestInfo = append([]byte("pong"))

    }

    if DEBUG == 1 {
        log.Println("==> guestInfo: ", string(guestInfo))
    }

    virtiorw.Write(guestInfo)
    virtiorw.Flush()
    virtiorw.Write([]byte(""))
    virtiorw.Close()
}

func readConsole(vconsole string) {
    virtiord, err := term.Open(vconsole, term.Speed(9600), term.RawMode)
    plugins.CheckErr(err)
    virtiord.Flush()

    buffer := make([]byte, 128)
    for {
        for {
            nlines, err := virtiord.Read(buffer)
            plugins.CheckErr(err)

            if DEBUG == 1 {
                log.Println("Virtio Received: ", string(buffer[:nlines]))
            }

            go writeConsole(string(buffer[:nlines]))
        }
        virtiord.Flush()
    }
    virtiord.Close()
}

func main() {
    vconsole := checkConsole()
    readConsole(vconsole)
}
