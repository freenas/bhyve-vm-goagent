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

	fmt.Printf("==> Running server at: %s:%d\n", *ipaddr, *port)

	log.Fatal(http.ListenAndServe(bindto, nil))
}
