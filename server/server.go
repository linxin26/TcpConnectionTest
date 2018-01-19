/**
@author linx
@date 2018-01-19
 */
package main

import (
	"net"
	"strings"
	"time"
	"strconv"
	"sync"
	"sync/atomic"
	"flag"
	"fmt"
)

var m sync.Map
var sum int32

func main() {

	bindAddr := flag.String("bindAddr", "0.0.0.0:8009", "bindAddr")
	flag.Parse()

	listener, error := net.Listen("tcp", *bindAddr)

	if error != nil {
		println(error.Error())
		listener.Close()
	}

	println("bind: ", *bindAddr)

	go print()

	for {
		conn, err := listener.Accept()
		if err == nil {
			go processConn(conn)
		} else {
			println(err.Error())
		}
	}

}

func processConn(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	info := strings.Split(addr, ":")
	port, _ := strconv.Atoi(info[1])
	atomic.AddInt32(&sum, 1)
	m.Store(addr, port)
	//go tickerSend(addr, conn)
}

func tickerSend(data string, conn net.Conn) {
	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			//_, err := conn.Write([]byte(data))
			//if err != nil {
			//	println(err)
			//}
		}
	}
}

func print() {
	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			total := 0
			m.Range(func(k, v interface{}) bool {
				total++
				return true
			})
			fmt.Printf("当前连接总数:%d \n", total)
		}
	}
}
