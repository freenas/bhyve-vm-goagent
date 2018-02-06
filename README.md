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

It will try to detect your OS and build the binaries for it.

## Build release
1) `make deps`
2) `make deps_windows`
3) `make release`

It will cross compile binaries for the following platforms: FreeBSD, NetBSD, Linux and Windows.
NOTE: When it runs `make deps_windows` a 'build constraints' will appears, just ignore it.

## Usage instructions (Host VM)
Start bhyve(8).
```
Usage of -virtio:bhyve -A -H -w -c 2 -m 2048 -s 0:0,hostbridge \
-s 31,lpc -l com1,/dev/nmdm132A -l bootrom,/usr/local/share/uefi-firmware/BHYVE_UEFI.fd \
-s 3,e1000,tap0,mac=00:a0:98:1e:c0:08 -s 9,fbuf,vncserver,tcp=192.168.100.111:6032,w=800,h=600,, \
-s 30,xhci,tablet -s 4,ahci-hd,/dev/zvol/tank/freebsd12 \
-s 2,virtio-console,org.freenas.bhyve-agent=/tmp/FreeBSD12.sock FreeBSD12
```

Note for the parameter <b>-s 2,virtio-console,org.freenas.bhyve-agent=/tmp/FreeBSD12.sock</b>

## Usage instructions (Guest VM)
`root@guest:/tmp # ./bhyve-vm-goagent-freebsd-386`
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

<b> Start in websocket mode</b>
```
root@guest:/tmp # ./bhyve-vm-goagent-freebsd-386 -websocket -ipaddr="192.168.100.193" -port=9191
018/01/22 11:30:12 ==> Running server at: 192.168.100.193:9191
```
<b> Start in virtio mode</b>
```
root@guest:/tmp # ./bhyve-vm-goagent-freebsd-386 -virtio
2018/01/22 11:41:50 ==> /dev/vtcon/org.freenas.bhyve-agent
```

## How to use host [tools](https://github.com/araujobsd/bhyve-vm-goagent/tree/master/tools/host)?
<b> goserial.go:</b>
```
root@freenas:/tmp # ./host -socket="/tmp/FreeBSD12.sock" -ether -uptime
Result:  [{"mtu":1500,"name":"em0","hardwareaddr":"00:a0:98:1e:c0:08","flags":["up","broadcast","multicast"],"addrs":[{"addr":"192.168.100.193/24"}]} {"mtu":16384,"name":"lo0","hardwareaddr":"","flags":["up","loopback","multicast"],"addrs":[{"addr":"::1/128"},{"addr":"fe80::1/64"},{"addr":"127.0.0.1/8"}]}]
Result:  {"seconds":5742,"days":0,"hours":1,"minutes":35}
```

<b> client.py:</b>
```
[araujo@pipoca] /z/go/bhyve-vm-goagent/tools/host# python3 client.py
```

You need to edit client.py and change GUEST_IP and GUEST_PORT according with your guest vm settings.


## Copyright and licensing
Distributed under [2-Clause BSD License](https://github.com/araujobsd/bhyve-vm-goagent/blob/master/LICENSE).
