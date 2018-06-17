package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"syscall"

	"github.com/mushroomsir/go-examples/util"
	"github.com/mushroomsir/logger/alog"
)

var (
	iface = flag.String("iface", "eth0", "net interface name")
)

func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

// https://github.com/spotify/linux/blob/master/include/linux/if_ether.h
// http://man7.org/linux/man-pages/man7/packet.7.html
func main() {
	flag.Parse()
	ifi, err := net.InterfaceByName(*iface)
	util.CheckError(err)
	alog.Info(ifi)
	// 2048 ipv4
	// 52320  custom
	// 拷贝一份
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_DGRAM, int(htons(3)))
	util.CheckError(err)
	// syscall.Bind(fd, &syscall.SockaddrLinklayer{
	// 	Protocol: 2048,
	// 	Ifindex:  ifi.Index,
	// })
	for {
		buf := make([]byte, 1518)
		n, from, err := syscall.Recvfrom(fd, buf, 0)
		util.CheckError(err)
		if n < 14 {
			//	continue
		}
		header := ParseHeader(buf[0:14])
		fmt.Println(header, from)
	}
}

// Header ...
type Header struct {
	DestinationAddress []byte
	SourceAddress      []byte
	EtherType          uint16
	//FCS  以太网帧尾部的FCS，发送的时候由硬件计算并添加，接收的时候由硬件校验并去除。
}

// ParseHeader ...
func ParseHeader(buf []byte) *Header {
	header := new(Header)
	header.DestinationAddress = buf[0:6]
	header.SourceAddress = buf[6:12]
	header.EtherType = binary.BigEndian.Uint16(buf[12:14])
	return header
}
