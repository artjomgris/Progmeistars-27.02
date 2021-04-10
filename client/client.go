package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var data struct {
		X int32
		Y int32
		N int32
	}

	var i int32 = 0
	var j int32 = 12
	for ; i <= 12; i++ {
		data.X = i
		data.Y = i
		data.N = i
		var buf bytes.Buffer
		err = binary.Write(&buf, binary.LittleEndian, data)
		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}
	i -= 2
	for ; i >= 1; i-- {
		data.X = 12 + (12 - i)
		data.Y = i
		data.N = j
		var buf bytes.Buffer
		err = binary.Write(&buf, binary.LittleEndian, data)
		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
		j++
	}

	for ; i <= 12; i++ {
		data.X = 12 + i
		data.Y = i
		data.N = i
		var buf bytes.Buffer
		err = binary.Write(&buf, binary.LittleEndian, data)
		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}
}
