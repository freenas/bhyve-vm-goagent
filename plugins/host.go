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

// Return pong.
func Ping() string {
	return "pong"
}
