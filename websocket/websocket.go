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

package websocket

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/araujobsd/bhyve-vm-goagent/plugins"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func testIPAddr(ipadd string) bool {
	ipType := net.ParseIP(ipadd)
	if ipType.To4() == nil || ipType.To16() == nil {
		return false
	}
	return true
}

func home(w http.ResponseWriter, r *http.Request) {
	var dump []byte
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Upgrade: ", err)
		return
	}

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}

		switch opt := string(message); opt {
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
		err = c.WriteMessage(mt, dump)
		plugins.CheckErr(err)
	}
	defer c.Close()
}

func RunServer(ipaddr *string, port *int) {
	var bindto string
	http.HandleFunc("/", home)

	if testIPAddr(*ipaddr) {
		bindto = *ipaddr + ":" + strconv.Itoa(*port)
	} else {
		bindto = "127.0.0.1:" + strconv.Itoa(*port)
	}

	log.Printf("==> Running server at: %s:%d\n", *ipaddr, *port)
	log.Fatal(http.ListenAndServe(bindto, nil))
}
