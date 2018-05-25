package main

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/mushroomsir/go-examples/util"
)

func main() {
	go func() {
		netaddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
		conn, err := net.ListenIP("ip4:tcp", netaddr)
		util.CheckError(err)
		for {
			buf := make([]byte, 1024)
			n, _, err := conn.ReadFrom(buf)
			util.CheckError(err)
			tcp := NewTCPHeader(buf[0:n])
			log.Println(n,"listen: ", tcp)
		}
	}()
	time.Sleep(time.Second)
	send("172.17.0.2", "127.0.0.1")
	time.Sleep(time.Second)
}
func send(local string, remote string) {
	packet := TCPHeader{
		Source:      17663, // Random ephemeral port
		Destination: 8020,
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

	conn, err := net.Dial("ip4:tcp", remote)
	if err != nil {
		log.Fatalf("Dial: %s\n", err)
	}
	conn.Write(data)
	conn.Write(data)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	util.CheckError(err)
	///tcp := NewTCPHeader(buf[0:n])
	log.Println(n ," send:", buf[0:n])
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

// ipheader := &ipv4.Header{
// 	Version:  4,
// 	Len:      20,
// 	TOS:      1,
// 	TotalLen: 1,
// 	ID:       0xcafe,
// 	Flags:    ipv4.DontFragment,
// 	FragOff:  1500,
// 	TTL:      255,
// 	Protocol: 1,
// 	Checksum: 0xdead,
// 	Src:      net.IPv4(172, 16, 254, 254),
// 	Dst:      net.IPv4(127, 0, 0, 1),
// }
// ipheadByte, err := ipheader.Marshal()
// _=ipheadByte
// util.CheckError(err)
// //ipheadByte = append(ipheadByte, []byte{2}...)

// //conn.Write(ipheadByte)
// //conn.Write([]byte("x"))
