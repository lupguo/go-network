package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tkstorm/go-network/tcp/internal/helper"
	"net"
	"os"
	"os/signal"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", "127.0.0.1:6666", "connect network address")
	flag.Parse()
}

func main() {
	// dial server
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		log.Fatalln("server dial fail", err)
	}

	log.Println("client dial to server", addr)
	// go read from conn
	wait := make(chan struct{})
	recvCh := make(chan struct{})
	go func() {
		ReceiveAsciiMessage(conn, recvCh)
		wait <- struct{}{}
	}()

	// get message from stdin
	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt)
	go InputAsciiMessage(conn, signCh, recvCh)

	// wait srv send
	<-wait
}

//InputAsciiMessage
func InputAsciiMessage(conn net.Conn, shutdown <-chan os.Signal, recvCh chan<- struct{}) {
	var content string
	for {
		select {
		case <-shutdown:
			recvCh <- struct{}{}
			return
		default:
			log.Println("input message >>>")
			if _, err := fmt.Scanln(&content); err != nil {
				continue
			}
			if _, err := fmt.Fprintln(conn, content); err != nil {
				log.Println("input error,", err)
			}
		}
	}
	conn.Close()
}

//ReceiveAsciiMessage
func ReceiveAsciiMessage(conn net.Conn, out <-chan struct{}) {
	var content string
	for {
		select {
		case <-out:
			return
		default:
			log.Println("begin receive message <<")
			if _, err := fmt.Scanln(&content); err != nil {
				log.Println("recv error:", err)
				continue
			}
			helper.Prompt(conn.RemoteAddr().String(), helper.NewMessage(len(content), []byte(content)))
		}
		//time.Sleep(100 * time.Millisecond)
	}
	conn.Close()
}

//ReceiveGobMessage
func ReceiveGobMessage(conn net.Conn, wait chan<- struct{}) {
	var replayMessage string
	if err := gob.NewDecoder(conn).Decode(&replayMessage); err != nil {
		log.Println("client receive err", err)
	}
	log.Println("client receive", replayMessage)

	wait <- struct{}{}
}
