package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/nsf/termbox-go"
	"net"
)

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
	var n int32 = 0
	ch := make(chan point)
	var p point
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	for {
		go handleConnection(listener, ch)
		p = <-ch
		if p.N >= n {
			termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(int(p.X), int(p.Y), '*', termbox.ColorRed, termbox.ColorBlack)
			termbox.SetCell(int(p.X-1), int(p.Y), '*', termbox.ColorGreen, termbox.ColorBlack)
			termbox.SetCell(int(p.X+1), int(p.Y), '*', termbox.ColorGreen, termbox.ColorBlack)
			termbox.SetCell(int(p.X), int(p.Y+1), '*', termbox.ColorGreen, termbox.ColorBlack)
			termbox.SetCell(int(p.X), int(p.Y-1), '*', termbox.ColorGreen, termbox.ColorBlack)
			termbox.SetCell(int(p.X+1), int(p.Y+1), '*', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(int(p.X-1), int(p.Y+1), '*', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(int(p.X+1), int(p.Y-1), '*', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(int(p.X-1), int(p.Y-1), '*', termbox.ColorWhite, termbox.ColorBlack)
			termbox.Flush()
			n++
		}

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
