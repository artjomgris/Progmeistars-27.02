package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

var layout = "2006-01-02T15:04:05.000Z"

type point struct {
	X int32
	Y int32
	N int32
}

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
	//var n int32 = 0
	var ch chan point
	//termbox.Init()
	//	defer termbox.Close()
	for {
		//termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
		handleConnection(listener, ch)
		p := <-ch
		//termbox.SetCell(int(p.X), int(p.Y), '*', termbox.ColorRed, termbox.ColorWhite)
		fmt.Println(p)
		//termbox.Flush()
		time.Sleep(2 * time.Second)
	}

}

func handleConnection(con *net.UDPConn, ch chan point) {
	buf := make([]byte, 2000)
	n, err := con.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n])

	var data point

	err = binary.Read(buff, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- data
}
