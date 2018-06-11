package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/mushroomsir/go-examples/util"
	"github.com/mushroomsir/logger/alog"
)

var (
	iface = flag.String("iface", "eth0", "net interface name")
)

func main() {
	flag.Parse()
	ifi, err := net.InterfaceByName(*iface)
	util.CheckError(err)
	alog.Info(ifi)
	fd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 52428)
	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))
	for {
		buf := make([]byte, 1518)
		n, err := f.Read(buf)
		util.CheckError(err)
		fmt.Println(buf[0:n])
	}
}
