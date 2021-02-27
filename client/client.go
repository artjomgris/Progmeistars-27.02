package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "localhost:5000")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data struct {
		X    int32
		Y    int32
		Time []byte
	}

	data.X = 1
	data.Y = 2
	data.Time = []byte(time.Now().String())

	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, data)

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
}
