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
	adr, err := net.ResolveUDPAddr("udp", "localhost:5000")
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
	buf := make([]byte, 3000)
	n, err := con.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n])

	var data struct {
		X    int32
		Y    int32
		Time []byte
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
