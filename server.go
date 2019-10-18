package main

import (
	"os/signal"
	"flag"
	"fmt"
	"net"
	"dbs"
	"strconv"
	"sync/atomic"
	"syscall"
)

func main() {
	signal.Ignore(syscall.SIGPIPE)
	host := flag.String("host", "localhost:8888", "IP:PORT")
	flag.Parse()
	ln, err := net.Listen("tcp", *host)
	defer ln.Close()
	if err != nil {
		fmt.Println(err)
	}
	cnt := new(uint32)
	*cnt = 0
	done := false
	capacity := 500
	ch := make(chan dbs.Pack ,capacity)
	ok := make(chan bool)
	dbaddr := "mongodb://localhost:27017"
			defer close(ch)
			defer close(ok)
	for {
		conn, err := ln.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("connection successful")
		go handle(conn,ch,ok,cnt)
		if(!done){
		go dbs.Storage(dbaddr, ch, capacity ,ok)
		done = true
		}
	}
}

func handle(conn net.Conn,ch chan dbs.Pack,ok chan bool,cnt *uint32) {
	defer conn.Close()
		userid := make([]byte,7);
	conn.Read(userid)
	id , _ := strconv.Atoi(string(userid))
	fmt.Printf("Client No.%d connect\n",id)
	defer fmt.Printf("Client No.%d disconnect\n",id)
	var pkg dbs.Pack
	for{		sumln := 0
				buf := make([]byte,512)
				ln ,err :=	conn.Read(buf);
				if err ==nil && ln != 0 {

						pkg = dbs.Pack{Id : id , Comment:string(buf[sumln:sumln+ln])}
						sumln += ln
						if sumln >= 511{ pkg =dbs.Pack{Id: id , Comment:string(buf[sumln-ln:511])}
										sumln = 0	}
						ch <- pkg
						atomic.AddUint32(cnt,1)
					if  atomic.LoadUint32(cnt) >= 500{
							ok<-true
							atomic.StoreUint32(cnt , 0)
					}
				}else {
						return
				}

		}
}
