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

// Example of usage:
// ./host -socket="/tmp/FreeBSD12.sock" -disk -uptime
// Result:  [{"mountpoint":"/","fstype":"zfs","total":48978477056,"free":48396845056,"used":581632000}]
// Result:  {"seconds":3822,"days":0,"hours":1,"minutes":3}

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var socketPath = flag.String("socket", "", "Canonical path for the socket file.")
var memoryinfo = flag.Bool("mem", false, "Get memory information from guest.")
var cpuinfo = flag.Bool("cpu", false, "Get CPU information from guest.")
var diskinfo = flag.Bool("disk", false, "Get DISK information from guest.")
var etherinfo = flag.Bool("ether", false, "Get Ethernet information from guest.")
var uptimeinfo = flag.Bool("uptime", false, "Get guest's UPTIME.")
var ping = flag.Bool("ping", false, "Check if guest is alive.")

func connect(opt string, path string) {
	var buf [2048]byte

	c, err := net.Dial("unix", path)
	if err != nil {
		panic(err)
	}

	_, err = c.Write([]byte(opt))
	if err != nil {
		log.Println("Erro to write: ", err)
		panic(err)
	}

	n, err := c.Read(buf[:])
	if err != nil {
		log.Println("Erro to read: ", err)
		panic(err)
	}
	fmt.Println("Result: ", string(buf[:n]))

	c.Close()
}

func main() {
	var arg string
	flag.Parse()

	if len(*socketPath) > 0 {
		if *memoryinfo {
			arg = "mem"
			connect(arg, *socketPath)
		}
		if *cpuinfo {
			arg = "cpu"
			connect(arg, *socketPath)
		}
		if *diskinfo {
			arg = "disk"
			connect(arg, *socketPath)
		}
		if *etherinfo {
			arg = "iface"
			connect(arg, *socketPath)
		}
		if *uptimeinfo {
			arg = "uptime"
			connect(arg, *socketPath)
		}
		if *ping {
			arg = "ping"
			connect(arg, *socketPath)
		}
	}
}
