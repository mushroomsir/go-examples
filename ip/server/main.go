package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"

	"github.com/mushroomsir/go-examples/util"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8020")
	util.CheckError(err)

	listen, err := net.ListenTCP("tcp", addr)
	util.CheckError(err)
	buf := make([]byte, 1024)
	for {
		conn, err := listen.AcceptTCP()
		util.CheckError(err)
		len, err := conn.Read(buf)
		util.CheckError(err)
		fmt.Printf("%v", string(buf[:len]))

		// send()
	}
}

//remote := "192.168.0.118"
//"192.168.0.118"
//local := "127.0.0.1"
//send(local, remote)
// netaddr, _ := net.ResolveIPAddr("ip4", local)
// 	conn, err := net.ListenIP("ip4:tcp", netaddr)
// 	util.CheckError(err)

// 	buf := make([]byte, 1024)
// 	numRead, _, err := conn.ReadFrom(buf)
// 	util.CheckError(err)
// 	fmt.Printf("%v", string(buf[:numRead]))
func send(local string, remote string) {
	packet := TCPHeader{
		Source:      0xaa47, // Random ephemeral port
		Destination: 80,
		SeqNum:      rand.Uint32(),
		AckNum:      0,
		DataOffset:  5,      // 4 bits
		Reserved:    0,      // 3 bits
		ECN:         0,      // 3 bits
		Ctrl:        2,      // 6 bits (000010, SYN bit set)
		Window:      0xaaaa, // size of your receive window
		Checksum:    0,      // Kernel will set this if it's 0
		Urgent:      0,
		Options:     []TCPOption{},
	}

	data := packet.Marshal()
	packet.Checksum = Csum(data, to4byte(local), to4byte(remote))
	data = packet.Marshal()

	conn, err := net.Dial("ip4:ip", remote)
	if err != nil {
		log.Fatalf("Dial: %s\n", err)
	}

	conn.Write(data)
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	util.CheckError(err)
	log.Println(buf[0:n])
}
func to4byte(addr string) [4]byte {
	parts := strings.Split(addr, ".")
	b0, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("to4byte: %s (latency works with IPv4 addresses only, but not IPv6!)\n", err)
	}
	b1, _ := strconv.Atoi(parts[1])
	b2, _ := strconv.Atoi(parts[2])
	b3, _ := strconv.Atoi(parts[3])
	return [4]byte{byte(b0), byte(b1), byte(b2), byte(b3)}
}
