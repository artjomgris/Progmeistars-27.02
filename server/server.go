package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

var layout = "2006-01-02T15:04:05.000Z"

func main() {
	adr, err := net.ResolveUDPAddr("udp", "localhost:12400")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		handleConnection(listener)
	}

}

func handleConnection(con *net.UDPConn) {
	buf := make([]byte, 10000)
	n, err := con.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n])

	var data struct {
		X    int32
		Y    int32
		Time []uint8
	}

	err = binary.Read(buff, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := time.Parse(layout, string(data.Time))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data, t)
}
