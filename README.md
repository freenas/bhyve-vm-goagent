[![GoDoc](https://godoc.org/github.com/araujobsd/bhyve-vm-goagent/plugins?status.svg)](https://godoc.org/github.com/araujobsd/bhyve-vm-goagent/)
[![GitHub issues](https://img.shields.io/github/issues/araujobsd/bhyve-vm-goagent.svg)](https://github.com/araujobsd/bhyve-vm-goagent/issues)
[![GitHub forks](https://img.shields.io/github/forks/araujobsd/bhyve-vm-goagent.svg)](https://github.com/araujobsd/bhyve-vm-goagent/network)

bhyve-vm-goagent
================
It is an agent developed using [Go](http://golang.org/) that runs inside guest virtual machines and allows host hypervisor to obtain information from guest such like memory usage, cpu, disk and ethernet configuration.

It supports [virtio-console](https://fedoraproject.org/wiki/Features/VirtioSerial) as well as [WebSocket](http://www.rfc-editor.org/rfc/rfc6455.txt) protocol.

## Build instructions
1) `make deps`
2) `make`

It will cross compile binaries for the following platforms: FreeBSD, NetBSD, Linux and Windows.

## Usage instructions
`root@freenas:/tmp # ./bhyve-vm-goagent-freebsd-386`
```
Usage of -virtio:
  -virtio
	It uses the virtio-console for communication.

Usage of -websocket:
  -websocket
	It uses the websocket for communication.
  -ipaddr string
	IPv4/IPv6 to bind the service. (default "127.0.0.1")
  -port int
	TCP port to bind the websocket. (default 8080)
```

## Copyright and licensing
Distributed under 2-Clause BSD License.
