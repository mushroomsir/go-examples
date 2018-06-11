package main

import (
	"encoding/binary"
	"flag"
	"net"
	"syscall"
	"time"

	"github.com/mushroomsir/go-examples/util"
	"github.com/mushroomsir/logger/alog"
)

var (
	iface     = flag.String("iface", "eth0", "net interface name")
	broadcast = net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
)

func main() {
	flag.Parse()
	ifi, err := net.InterfaceByName(*iface)
	util.CheckError(err)
	alog.Info(ifi)
	fd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 52428)
	for {
		payload := []byte("msg")
		minPayload := len(payload)
		if minPayload < 46 {
			minPayload = 46
		}
		b := make([]byte, 14+minPayload)
		copy(b[0:6], broadcast)
		copy(b[6:12], ifi.HardwareAddr)

		etype := make([]byte, 2)
		binary.BigEndian.PutUint16(etype, uint16(52428))
		copy(b[12:14], etype)
		copy(b[14:14+len(payload)], payload)

		var baddr [8]byte
		copy(baddr[:], broadcast)
		to := &syscall.SockaddrLinklayer{
			Ifindex:  ifi.Index,
			Halen:    uint8(len(broadcast)),
			Addr:     baddr,
			Protocol: 52428,
		}
		err = syscall.Sendto(fd, b, 0, to)
		util.CheckError(err)
		time.Sleep(time.Second)
	}
}
