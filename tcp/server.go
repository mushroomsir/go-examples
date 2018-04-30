package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"syscall"

	"github.com/mushroomsir/go-examples/util"
	"github.com/mushroomsir/logger/alog"
)

func main() {
	ln, _ := net.Listen("tcp", ":8081")
	fmt.Println(ln.Addr())
	for {
		conn, err := ln.Accept()
		if !alog.Check(err) {
			go handleConn(conn)
		}
	}
}

const (
	keepalive = 1
)

func handleConn(conn net.Conn) {
	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		return
	}
	file, err := tcpconn.File()
	util.CheckError(err)

	err = syscall.SetsockoptInt(syscall.Handle(file.Fd()), syscall.SOL_SOCKET, syscall.SO_KEEPALIVE, keepalive)
	file.Close()
	util.CheckError(err)

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Received:", string(message))
	newmessage := strings.ToUpper(message)
	conn.Write([]byte(newmessage + "\n"))
}
