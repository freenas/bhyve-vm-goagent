bhyve-vm-goagent
================
It is an agent developed using Go that runs inside guest virtual machines and allows host hypervisor to obtain information from guest such like memory usage, cpu, disk and ethernet configuration.

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
