package main

import (
	"net"
	"fmt"
	"flag"
	"time"
	"sync/atomic"
)

/**
docker run -it --name client2 solinx.co/market/demo-client:0.1 client -serverAddr=172.16.187.228:28009 -localAddr=127.0.0.1
docker run -it --rm --name client3 --privileged solinx.co/market/demo-client:0.1 sh -c "sysctl -p && client -serverAddr=172.16.187.228:28009 -total=20000"
 */

var clientSum int32

func main() {

	serverAddr := flag.String("serverAddr", "127.0.0.1:8009", "serverAddr")
	localAddr := flag.String("localAddr", "127.0.0.1", "localAddr")
	total:=flag.Int("total",2,"conn total")

	flag.Parse()
	i := 0
	var cnum chan int
	go printSummary()
	for i = 0; i < *total; i++ {
		time.Sleep(time.Millisecond*2)
		 go openConn(*localAddr,*serverAddr)
	}

	<-cnum
}

func openConn(localAddr,serverAddr string){

	var cnum chan int
	conn, _err :=net.Dial("tcp",serverAddr)
	if _err != nil {
		fmt.Println(_err)
	}
	atomic.AddInt32(&clientSum, 1)
	//reader := bufio.NewReader(conn)
	//for {
	//	buf := make([]byte, 128)
	//	reader.Read(buf)
	//	if _err != nil {
	//		println(_err)
	//	}
	//	println(string(buf))
	//}

	defer func() {
		if conn != nil {
			println("关闭",conn.LocalAddr().String())
			conn.Close()
		}
	}()

	<-cnum

}

func printSummary() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("当前请求连接总数:%d \n",atomic.LoadInt32(&clientSum))
		}
	}
}
