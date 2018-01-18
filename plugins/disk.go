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

	"github.com/shirou/gopsutil/disk"
)

var FSTYPE = map[string]bool{
	"ext3": true,
	"ext4": true,
	"zfs":  true,
	"ufs":  true,
}

type DiskUsage struct {
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Total      uint64 `json:"total"`
	Free       uint64 `json:"free"`
	Used       uint64 `json:"used"`
}

var DiskInformation []DiskUsage

func UsageInfo(path string) DiskUsage {
	usage, err := disk.Usage(path)
	CheckErr(err)

	diskusage := DiskUsage{}
	diskusage.Mountpoint = usage.Path
	diskusage.Fstype = usage.Fstype
	diskusage.Total = usage.Total
	diskusage.Free = usage.Free
	diskusage.Used = usage.Used

	return diskusage
}

// Returns information about all slices.
func DiskInfo() []byte {
	vdisk, err := disk.Partitions(true)
	DiskInformation := []DiskUsage{}
	CheckErr(err)
	for _, disk := range vdisk {
		if FSTYPE[disk.Fstype] {
			info := UsageInfo("/")
			DiskInformation = append(DiskInformation, info)
			break
		}
	}
	convjson, err := json.Marshal([]DiskUsage(DiskInformation))
	CheckErr(err)

	return convjson
}
