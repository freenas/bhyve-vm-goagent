/*-
 * Copyright 2018 iXsystems, Inc.
 * Copyright 2018 by Marcelo Araujo <araujo@ixsystems.com>.
 * All rights reserved
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted providing that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
 * IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/araujobsd/bhyve-vm-goagent/plugins"
	"github.com/araujobsd/bhyve-vm-goagent/termios"
	"github.com/araujobsd/bhyve-vm-goagent/websocket"
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

func Run(fd int) {
	var dump []byte
	callback := termios.Read(fd)
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
		case "ping":
			pinginfo := plugins.Ping()
			dump = append([]byte(fmt.Sprintf("%v", pinginfo)))
		default:
			dump = append([]byte("bang! bang!"))
		}
		termios.Write(fd, dump)
	}
}

func usage() {
	fmt.Println("Usage of -virtio:")
	fmt.Printf("  -virtio\n")
	fmt.Printf("\tIt uses the virtio-console for communication.\n")
	fmt.Println("")
	fmt.Println("Usage of -websocket:")
	fmt.Printf("  -websocket\n")
	fmt.Printf("\tIt uses the websocket for communication.\n")
	fmt.Printf("  -ipaddr string\n")
	fmt.Printf("\tIPv4/IPv6 to bind the service. (default '127.0.0.1')\n")
	fmt.Printf("  -port int\n")
	fmt.Printf("\tTCP port to bind the websocket. (default 8080)\n")
	fmt.Println("")
	os.Exit(2)
}

func main() {
	vconsoleCmd := flag.NewFlagSet("-virtio", flag.ExitOnError)
	websocketCmd := flag.NewFlagSet("-websocket", flag.ExitOnError)
	websocketIp := websocketCmd.String("ipaddr", "127.0.0.1", "IPv4/IPv6 to bind the service.")
	websocketPort := websocketCmd.Int("port", 8080, "TCP port to bind the websocket.")

	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "-virtio":
		vconsoleCmd.Parse(os.Args[2:])
	case "-websocket":
		websocketCmd.Parse(os.Args[2:])
	default:
		usage()
	}

	if vconsoleCmd.Parsed() {
		vconsole := checkConsole()
		if len(vconsole) > 0 {
			fd := termios.NewConnection(vconsole)
			for {
				Run(fd)
			}
			defer termios.CloseConnection(fd)
		} else {
			fmt.Println("Error: Guest without support of virtio-console.")
			os.Exit(2)
		}
	} else if websocketCmd.Parsed() {
		websocket.RunServer(websocketIp, websocketPort)
	}
}
